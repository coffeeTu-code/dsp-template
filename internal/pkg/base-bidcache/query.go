package base_bidcache

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"dsp-template/api/base"
	helper_crypto "dsp-template/pkg2/helper-crypto"
)

type QueryParser interface {
	ParamOK() bool
	GetExchange() string
	GetExchangeConfig() *ExchangeConfig
	GetDeviceID() string
	GetPlacementID() string
	GetAlgoInfo() string
	SetAlgoInfo(string)
	GetBidfloor() float64
	GetBidInfo() *CacheValue
	GetAdtype() string
}

func BuildQuery(feature *base.Feature, adxCfg *ExchangeConfig) (*QueryInfo, error) {
	if feature.DeviceIds == nil || feature.App == nil || adxCfg == nil {
		return nil, errors.New("")
	}
	return &QueryInfo{
		QueryParser: BuildRawQuery(feature, adxCfg),
	}, nil
}

type QueryInfo struct {
	cfg *Config
	QueryParser
	exchangeConfig *ExchangeConfig
	CacheBidInfo   *CacheValue
	bidInfo        *CacheValue

	passWhitelist int // 0,初始状态，1.通过。2.不过
	paramOK       int // 0.初始状态，1.通过，2.不过
	entryAbtest   bool
	blocked       bool

	entryF2xxAbtest bool

	overload float64

	cacheKey    string
	deviceId    string
	placementID string
	flowMark    string
}

type CacheValue struct {
	Price        float64 // 美元千次
	NonbidReason string
	Version      string
	Data         string
}

// 先简单解析
func NewCacheValue(str string) (value *CacheValue, err error) {
	value = new(CacheValue)
	seg := strings.Split(str, "\t")
	value.Price, err = strconv.ParseFloat(seg[0], 64)
	if len(seg) > 1 {
		value.Version = seg[1]
	}
	if len(seg) > 2 {
		value.NonbidReason = seg[2]
	}
	if len(seg) > 3 {
		value.Data = seg[3]
	}
	return
}

// [price, nbr, version, data] separate by "\t"
func (this *CacheValue) ToString() string {
	if this == nil {
		return "0\t0\t" + strconv.FormatInt(time.Now().Unix(), 10) + "\t"
	}
	return strconv.FormatFloat(this.Price, 'f', -1, 64) +
		"\t" + this.NonbidReason +
		"\t" + this.Version +
		"\t" + this.Data
}

// 解析request
func (this *QueryInfo) ParamOK() bool {
	if this.paramOK == 0 {
		if this.QueryParser == nil {
			this.paramOK = 2
		} else if !this.QueryParser.ParamOK() {
			this.paramOK = 2
		} else if this.GetExchangeConfig() == nil {
			this.paramOK = 2
		} else if this.GetPlacementID() == "" {
			this.paramOK = 2
		} else if this.GetDeviceID() == "" || !CheckDeviceId(this.GetDeviceID()) {
			this.paramOK = 2
		} else {
			this.paramOK = 1
		}
	}
	return this.paramOK == 1
}

func (this *QueryInfo) PassWhitelist() bool {
	if this.passWhitelist == 0 {
		if this.GetExchangeConfig() == nil {
			this.passWhitelist = 2
			return false
		}
		// by placementid
		this.passWhitelist = 2
		if len(this.GetExchangeConfig().PlacementID) > 0 {
			for _, id := range this.GetExchangeConfig().PlacementID {
				if id == this.GetPlacementID() {
					this.passWhitelist = 1
					break
				}
			}
		} else {
			this.passWhitelist = 1
		}
	}
	return this.passWhitelist == 1
}

// 判断是否进行abtest
func (this *QueryInfo) EntryAbtest() bool {
	if this.GetExchangeConfig() == nil {
		return false
	}
	// judge by ratio
	this.entryAbtest = false
	ratio := this.GetExchangeConfig().Ratio
	if ratio == 1 {
		this.entryAbtest = true
	}
	if ratio > 0 && ratio > RandByDeviceID(this.GetDeviceID()+this.GetPlacementID(), "THEMIS_DEVICE") {
		this.entryAbtest = true
	}
	return this.entryAbtest
}

// true：应该过滤，当命中缓存并且缓存出价小于请求底价
// false: 不过滤，没有命中缓存，或者缓存出价不小于请求底价
func (this *QueryInfo) ShouldBlock() bool {
	return this.CacheBidInfo != nil && (this.CacheBidInfo.Price*this.GetOverload() < this.GetBidfloor() || this.CacheBidInfo.Price == 0)
}

func (this *QueryInfo) SetBlock() {
	this.blocked = true
	if (this.GetExchangeConfig() != nil && this.GetExchangeConfig().CacheAsInfo) ||
		(this.cfg.Default != nil && this.cfg.Default.CacheAsInfo) {
		if this.CacheBidInfo != nil {
			this.SetAlgoInfo(this.CacheBidInfo.Data)
		}
	}
}

// 返回true，表示应该放行
// 返回false，表示应该拦截
func (this *QueryInfo) EntryF2xxTest() bool {
	if this.GetExchangeConfig() != nil && this.GetExchangeConfig().UseF2xxAbtest {
		if RandByDeviceID(this.GetDeviceID(), "THEMIS_DEVICE_F2XX") < 0.5 {
			this.entryF2xxAbtest = true
		}
	}
	return this.entryF2xxAbtest
}

func (this *QueryInfo) GetOverload() float64 {
	if this.overload == 0 {
		for _, exchangeConfig := range []*ExchangeConfig{this.GetExchangeConfig(), this.cfg.Default} {
			if exchangeConfig == nil {
				continue
			}

			// 优先从condition拿，测试完对已经可以确定的adtype落地，扩量对剩下的继续abtest
			// 从condition_overload获取运行阶段的系数
			if this.overload == 0 && len(exchangeConfig.ConditionOverloads) != 0 {
				overload, ok := exchangeConfig.ConditionOverloads[this.GetOverloadCondition()]
				if !ok {
					overload, ok = exchangeConfig.ConditionOverloads["all"]
				}
				if ok {
					// overload_serial := gwutil.RandByDeviceID(this.GetDeviceID()+this.GetPlacementID(), "THEMIS_DEVICE_OVERLOAD")
					this.overload = overload.GetOverload(this.GetDeviceID() + this.GetPlacementID() + "THEMIS_DEVICE_OVERLOAD")
				}
			}

			// 从condition_overload获取运行阶段的系数
			if this.overload == 0 && len(exchangeConfig.ConditionOverload) != 0 {
				overload, ok := exchangeConfig.ConditionOverload[this.GetOverloadCondition()]
				if ok {
					this.overload = overload
				} else if overload, ok = exchangeConfig.ConditionOverload["all"]; ok {
					this.overload = overload
				}
			}

			// 从abtest_overload获取测试阶段的系数
			if this.overload == 0 && !exchangeConfig.AbTestOverload.IsEmpty() {
				// overload_serial := gwutil.RandByDeviceID(this.GetDeviceID()+this.GetPlacementID(), "THEMIS_DEVICE_OVERLOAD")
				this.overload = exchangeConfig.AbTestOverload.GetOverload(this.GetDeviceID() + this.GetPlacementID() + "THEMIS_DEVICE_OVERLOAD")
			}

			// 获取默认系数
			if this.overload == 0 && exchangeConfig.Overload != 0 {
				this.overload = exchangeConfig.Overload
			}
		}

		// 保底为1
		if this.overload == 0 {
			this.overload = 1
		}
	}
	return this.overload
}

func (this *QueryInfo) GetOverloadCondition() string {
	return this.QueryParser.GetAdtype()
}

// value为f1，表示需要进行缓存，但是没进入abtest
// value为f2，表示使用cache价格对请求进行了过滤，不更新缓存
// value为f3，表示使用cache价格进行了出价(暂时不支持)
// value为f4，表示向后转发了请求，使用了实时的出价
// value为f5，表示不需要进行缓存，在进入abtest之前判断
// value为f6，表示需要缓存但没命中白名单，在进入abtest之前判断
// abtest时候，主要是f1 vs f2+f4
func (this *QueryInfo) FlowMark() string {
	if this.flowMark == "" {
		if !this.ParamOK() {
			this.flowMark = "f5"
		} else if !this.PassWhitelist() {
			this.flowMark = "f6"
		} else if !this.entryAbtest {
			this.flowMark = "f1"
		} else if this.blocked {
			if this.GetExchangeConfig() != nil && this.GetExchangeConfig().UseF2xxAbtest {
				this.flowMark = "f2a"
			} else {
				this.flowMark = "f2"
			}
		} else {
			this.flowMark = "f4"
			if this.entryF2xxAbtest {
				this.flowMark = "f2b"
			}
		}
	}
	return this.flowMark
}

func (this *QueryInfo) GetCacheKey() string {
	if this.cacheKey == "" {
		this.cacheKey = this.GetExchange() + ":" + this.GetPlacementID() + ":" + this.GetDeviceID()
	}
	return this.cacheKey
}

func BuildRawQuery(feature *base.Feature, adxCfg *ExchangeConfig) *RawQueryParser {
	return &RawQueryParser{
		Googleadid:     feature.DeviceIds.Googleadid,
		Exchange:       feature.Exchange,
		Appid:          feature.App.PackageName,
		Adtype:         feature.AdType,
		ExchangeConfig: adxCfg,
		Bidfloor:       feature.BidFloor,
	}
}

type RawQueryParser struct {
	Googleadid     string
	Exchange       string
	Appid          string
	Adtype         string
	ExchangeConfig *ExchangeConfig
	DeviceID       string
	PlacementID    string
	Bidfloor       float64
	BidPrice       float64
	NonbidReason   int
}

func (this *RawQueryParser) GetAdtype() string {
	return this.Adtype
}

func (this *RawQueryParser) ParamOK() bool {
	return true
}

func (this *RawQueryParser) GetExchangeConfig() *ExchangeConfig {
	return this.ExchangeConfig
}

func (this *RawQueryParser) GetExchange() string {
	return this.Exchange
}

func (this *RawQueryParser) GetBidfloor() float64 {
	return this.Bidfloor
}

func (this *RawQueryParser) GetBidInfo() *CacheValue {
	return &CacheValue{
		Price:        this.BidPrice,
		NonbidReason: strconv.Itoa(this.NonbidReason),
		Version:      strconv.FormatInt(time.Now().Unix(), 10),
	}
}

func (this *RawQueryParser) GetPlacementID() string {
	if this.PlacementID == "" {
		switch this.Exchange {
		default:
			this.PlacementID = helper_crypto.Md5(this.GetExchange() + "_" + this.Adtype + "_" + this.Appid)
		}
		// 如果处理过以后还是为空则打上特殊标记
		if this.PlacementID == "" {
			this.PlacementID = "-"
		}
	}
	if this.PlacementID == "-" {
		return ""
	}
	return this.PlacementID
}

func (this *RawQueryParser) GetDeviceID() string {
	if this.DeviceID == "" {
		switch this.Exchange {
		default:
			this.DeviceID = this.Googleadid
		}
		// 如果处理完还是不合法，则打上特殊标记
		if this.DeviceID == "" {
			this.DeviceID = "-"
		}
	}
	if this.DeviceID == "-" {
		return ""
	}
	return this.DeviceID
}

func (this *RawQueryParser) GetAlgoInfo() string {
	return ""
}

func (this *RawQueryParser) SetAlgoInfo(string) {
}
