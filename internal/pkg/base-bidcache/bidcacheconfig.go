package base_bidcache

import (
	"hash/crc64"

	"dsp-template/api/backend"
)

type Config struct {
	Open         bool                       `toml:"open"`          // 是否开启缓存功能
	Region       string                     `toml:"region"`        // metrics的region, 如果需要实时获取，则留空并通过gwconst.InitRegion方法去初始化
	Default      *ExchangeConfig            `toml:"default"`       // 默认的exchange设置
	AdxConfig    map[string]*ExchangeConfig `toml:"adxconfig"`     // exchange设置
	Runtime      *RuntimeConfig             `toml:"runtime"`       // 实时更新的配置
	CacheDB      backend.RedisConfig        `toml:"cachedb"`       // cache的db配置
	RedisHystrix backend.RedisHystrixConfig `toml:"redis_hystrix"` // 使用hystrix来做熔断的配置
}

type (
	ExchangeConfig struct {
		// match_tag -> track_tag -> value
		Ratio              float64                    `toml:"ratio"`           // 进入实验组的比例
		Expire             int64                      `toml:"expire"`          // 缓存周期
		CacheAsInfo        bool                       `toml:"cache_asinfo"`    // 是否缓存额外内容
		UseF2xxAbtest      bool                       `toml:"use_f2xx_abtest"` // 是否开启二次abtest，默认50%对半分
		PlacementID        []string                   `toml:"placement_id"`    // pid白名单列表，非空有效
		Overload           float64                    `toml:"overload"`        // 底价判断的溢价系数
		AbTestOverload     OverloadValues             `toml:"abtest_overload"`
		ConditionOverloads map[string]*OverloadValues `toml:"condition_overloads"` // 设定abtest
		ConditionOverload  map[string]float64         `toml:"condition_overload"`  // 设定abtest
		BlockReserve       float64                    `toml:"block_reserve"`       // 缓存异常时的备用拦截比例
	}
	OverloadValues struct {
		Values  []float64     `toml:"values"`
		Section SectionConfig `toml:"section"`
	}
	SectionConfig struct {
		Min   float64 `toml:"min"`
		Max   float64 `toml:"max"`
		Slots int     `toml:"slots"`
	}
)

type RuntimeConfig struct {
	Exchanges map[string]*ExchangeConfig `toml:"exchanges"`
}

func (this *OverloadValues) IsEmpty() bool {
	return this.Section.Slots == 0 && len(this.Values) == 0
}

func (this *OverloadValues) GetOverload(id string) float64 {
	if len(this.Values) > 0 {
		return this.Values[PickByDevice(id, len(this.Values))]
	} else if this.Section.Slots != 0 {
		return float64(PickByDevice(id, this.Section.Slots))*(this.Section.Max-this.Section.Min)/float64(this.Section.Slots) + this.Section.Min
	}
	return 0
}

var crc_table = crc64.MakeTable(crc64.ECMA)

func RandByDeviceID(id string, salt string) float64 {
	return float64(crc64.Checksum([]byte(id+salt), crc_table)%1000) / 1000
}

func PickByDevice(id string, mod int) int {
	return int(crc64.Checksum([]byte(id), crc_table) % uint64(mod))
}

// 判断设备id是否合法
// 非法设备ID 如 deviceId == "" || deviceId == "-" || deviceId == "00000000-0000-0000-0000-000000000000"
// return bool 合法返回true, 非法返回false
func CheckDeviceId(deviceId string) bool {
	//tmpDeviceId := strings.TrimSpace(deviceId)
	//tmpDeviceId = strings.Replace(tmpDeviceId, "0", "", -1)
	//tmpDeviceId = strings.Replace(tmpDeviceId, "-", "", -1)
	for i := 0; i < len(deviceId); i++ {
		if deviceId[i] != '0' && deviceId[i] != '-' && deviceId[i] != ' ' {
			return true
		}
	}

	return false
}
