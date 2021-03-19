package base_bidcache

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-redis/redis"

	microservice_helper "dsp-template/pkg2/helper-microservice"
)

var defaultBidCache = NewBidCache()

func NewBidCache() *BidCache {
	return &BidCache{}
}

type BidCache struct {
	cacheDB *redis.ClusterClient
	cfg     *Config
}

func (this *BidCache) Init(cfg *Config) error {
	defaultBidCache.cfg = cfg

	if cfg.RedisHystrix.Use {
		hystrix.ConfigureCommand("Redis", hystrix.CommandConfig{
			Timeout:               cfg.RedisHystrix.Timeout,
			MaxConcurrentRequests: cfg.RedisHystrix.MaxConcurrentRequests,
			ErrorPercentThreshold: cfg.RedisHystrix.ErrorPercentThreshold,
			SleepWindow:           cfg.RedisHystrix.SleepWindow,
		})
	}

	vpath := strings.Split(cfg.CacheDB.Host, ",")
	this.cacheDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              vpath,
		Password:           cfg.CacheDB.Password,
		DialTimeout:        time.Duration(cfg.CacheDB.DialTimeout) * time.Millisecond,
		ReadTimeout:        time.Duration(cfg.CacheDB.ReadTimeout) * time.Millisecond,
		MaxRetries:         cfg.CacheDB.MaxRetries,
		ReadOnly:           cfg.CacheDB.ReadOnly,
		RouteRandomly:      cfg.CacheDB.RouteRandomly,
		PoolSize:           cfg.CacheDB.PoolSize,
		PoolTimeout:        time.Duration(cfg.CacheDB.PoolTimeout) * time.Millisecond,
		MinIdleConns:       cfg.CacheDB.MinIdleConns,
		MaxRedirects:       cfg.CacheDB.MaxRedirects,
		MaxConnAge:         time.Duration(cfg.CacheDB.MaxConnAge) * time.Second,
		IdleTimeout:        time.Duration(cfg.CacheDB.IdleTimeout) * time.Second,
		IdleCheckFrequency: time.Duration(cfg.CacheDB.IdleCheckFrequency) * time.Second,
		RouteByLatency:     cfg.CacheDB.RouteByLatency,
	})

	if res := this.cacheDB.Ping(); res.Err() != nil {
		return res.Err()
	} else if res.Val() != "PONG" {
		return errors.New("ping redis faild: " + cfg.CacheDB.Host)
	}

	return nil
}

func (this *BidCache) ResetConfig(cfg *Config) {
	defaultBidCache.cfg = cfg
}

// 竞价请求需要调用此函数访问缓存，
func (this *BidCache) Retrive(info *QueryInfo) (blocked bool, err error) {
	if !info.ParamOK() || !info.PassWhitelist() {
		return false, nil
	}

	if info.EntryAbtest() {
		// 首先查询cache
		hit, err := this.getCache(info)
		if err != nil {
			// 缓存读取失败的情况下，如果设置了备用比例，则使用备用比例作为命中率
			var reserve float64
			if info.GetExchangeConfig().BlockReserve != 0 {
				reserve = info.GetExchangeConfig().BlockReserve
			} else if this.cfg.Default.BlockReserve != 0 {
				reserve = this.cfg.Default.BlockReserve
			}
			if reserve != 0 && rand.Float64() < reserve {
				info.SetBlock()
			}
			return info.blocked, err
		}
		// 统计缓存命中情况
		// 如果命中缓存并且缓存出价低于底价，包含不出价的情形，则拦截
		if hit && info.ShouldBlock() {
			if !info.EntryF2xxTest() {
				info.SetBlock()
			}
		}
	}
	return info.blocked, nil
}

// 竞价响应前调用更新缓存
func (this *BidCache) Record(info *QueryInfo) error {
	if !info.ParamOK() || !info.PassWhitelist() {
		return nil
	}

	// f4 流量才需要写缓存,即命中ratio并且没有并底价block的:info.entryAbtest && !info.block
	if info.FlowMark() == "f4" || info.FlowMark() == "f2b" {
		return this.setCache(info)
	}
	return nil
}

func (this *BidCache) setCache(info *QueryInfo) error {
	key := info.GetCacheKey()
	value := info.GetBidInfo().ToString()
	var expire int64
	if info.GetExchangeConfig() != nil && info.GetExchangeConfig().Expire != 0 {
		expire = info.GetExchangeConfig().Expire
	}
	if expire == 0 && this.cfg.Default != nil && this.cfg.Default.Expire != 0 {
		expire = this.cfg.Default.Expire
	}
	if expire == 0 {
		expire = 5
	}

	var err error
	if this.cfg.RedisHystrix.Use {
		_, err = microservice_helper.CallDependentService("Redis", func() (interface{}, error) {
			return nil, this.setRedis(key, value, time.Second*time.Duration(expire))
		}, nil)
	} else {
		err = this.setRedis(key, value, time.Second*time.Duration(expire))
	}
	if err != nil {
		return err
	}
	return nil
}

// 是否命中缓存
func (this *BidCache) getCache(info *QueryInfo) (ok bool, err error) {
	key := info.GetCacheKey()

	var value string
	if this.cfg.RedisHystrix.Use {
		if ret, _err := microservice_helper.CallDependentService("Redis", func() (interface{}, error) {
			return this.getRedis(key)
		}, nil); _err != nil {
			err = _err
		} else if retStr, ok := ret.(string); ok {
			value = retStr
		}
	} else {
		value, err = this.getRedis(key)
	}
	if err != nil {
		return
	}

	if value != "" {
		if info.CacheBidInfo, err = NewCacheValue(value); err != nil { // 如果有并发访问info，需要小心覆盖
			return
		}
		ok = true
	} else {
	}
	return
}

func (this *BidCache) setRedis(key string, value string, expire time.Duration) (err error) {
	return this.cacheDB.Set(key, value, expire).Err()
}

func (this *BidCache) getRedis(key string) (value string, err error) {
	res := this.cacheDB.Get(key)
	if res.Err() == redis.Nil {
		return
	} else if res.Err() != nil {
		err = res.Err()
		return
	}
	return res.Val(), nil
}
