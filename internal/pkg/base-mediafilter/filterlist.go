package base_mediafilter

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"dsp-template/api/base"
	"dsp-template/api/dbstruct"
	base_bidcache "dsp-template/internal/pkg/base-bidcache"
)

/*

非法流量过滤

（1）流量字段缺失
（2）流量字段不合常规
（3）ip block
（4）low history roi
（5）bid cache

*/

// filter func 通用过滤逻辑
// return true 过滤流量
// return false 保留流量
type MediaFilter func(feature *dbstruct.Feature) bool

// options 自定义过滤条件
// return true 保留流量
// return false 过滤流量
type WhiteTableOption func(feature *dbstruct.Feature) bool

func reserved(feature *dbstruct.Feature, options ...WhiteTableOption) bool {
	for i := range options {
		if options[i](feature) {
			return true
		}
	}
	return false
}

func MediaHardFilterSize(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		if len(feature.Size) > 0 {
			return strings.ContainsAny(feature.Size[0], "-")
		}
		return false
	}
}

func MediaHardFilterTraffic(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		if len(feature.Size) > 0 && feature.Size[0] == "0x0" {
			switch feature.Exchange {
			case base.ExchangeOppo:
				return feature.AdType != base.AdTypeVideoRewarded && feature.AdType != base.AdTypeVideoInterstitial
			}
		}
		return false
	}
}

func MediaHardFilterCountry(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return len(feature.CountryCode) != 2
	}
}

func MediaHardFilterUA(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return feature.Device != nil && feature.Device.UserAgent == "" && !hasDeviceId(feature.DeviceIds)
	}
}

func hasDeviceId(deviceIds *dbstruct.FDeviceIds) bool {
	return deviceIds != nil &&
		(len(deviceIds.GoogleAdId) > 0 ||
			len(deviceIds.Idfa) > 0 ||
			len(deviceIds.GoogleAdIdMD5) > 0 ||
			len(deviceIds.GoogleAdIdSHA1) > 0 ||
			len(deviceIds.Imei) > 0 ||
			len(deviceIds.ImeiMD5) > 0 ||
			len(deviceIds.ImeiSHA1) > 0 ||
			len(deviceIds.AndroidId) > 0 ||
			len(deviceIds.AndroidIdMD5) > 0 ||
			len(deviceIds.AndroidIdSHA1) > 0 ||
			len(deviceIds.OAId) > 0 ||
			len(deviceIds.OAIdMD5) > 0)
}

func MediaHardFilterDeviceType(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return feature.Device != nil && isNonDeviceType(feature.Device)
	}
}

func isNonDeviceType(device *dbstruct.FDevice) bool {
	switch device.DeviceType {
	case base.DeviceTypeMobileTablet, base.DeviceTypePhone, base.DeviceTypeTablet:
		return false
	}
	return true
}

func MediaHardFilterBannerBType(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return feature.Imp != nil && isBlockImage(feature.Imp)
	}
}

func isBlockImage(imp *dbstruct.FImp) bool {
	for _, btype := range imp.BannerBType {
		if btype == base.BannerBTypeBANNER {
			return true
		}
	}
	return false
}

func MediaHardFilterPlatform(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return feature.Platform != base.PlatformAndroid && feature.Platform != base.PlatformIos
	}
}

func MediaHardFilterDeviceID(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return feature.DeviceIds != nil && isPseudoDeviceID(feature.DeviceIds) && rand.Float64() < pseudoRandom
	}
}

var (
	pseudoRandom = 0.8 //80%过滤  20%验流量
	googleIdExp  = regexp.MustCompile("[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}")
)

func isPseudoDeviceID(deviceIds *dbstruct.FDeviceIds) bool {
	if len(deviceIds.GoogleAdId) > 0 {
		return googleIdExp.MatchString(deviceIds.GoogleAdId)
	}
	return false
}

func MediaHardFilterBidCache(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return isBidCached(feature)
	}
}

func isBidCached(feature *dbstruct.Feature) bool {
	rawParse, err := base_bidcache.BuildQuery(feature, base_bidcache.ReadExchangeConfig(feature.Exchange))
	if err != nil {
		return false // 不过滤
	}
	block, err := base_bidcache.App().Retrive(rawParse)
	if err != nil {
		return false // 查询失败不过滤
	}
	return block
}

func MediaSoftFilterPackageName(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		if feature.App == nil || len(feature.App.PackageName) == 0 {
			return true
		}
		switch feature.Platform {
		case base.PlatformAndroid:
			return !isAndroidPackageName(feature.App.PackageName)
		case base.PlatformIos:
			return !isIosPackageName(feature.App.PackageName)
		default:
			return false // 非 app 类不过滤
		}
	}
}

func isIosPackageName(pkgName string) bool {
	_, err := strconv.Atoi(strings.TrimLeft(strings.ToLower(pkgName), "id"))
	return err == nil
}

func isAndroidPackageName(pkgName string) bool {
	pkgName = strings.ToLower(pkgName)
	return strings.HasPrefix(pkgName, "http") || strings.Contains(pkgName, "/") ||
		!strings.Contains(pkgName, ".")
}

func TrafficRelease(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return false
	}
}

func AntiCheat(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		switch {
		case feature.Device == nil || isBlackIP(feature.Device.IP):
			return true
		case feature.App == nil || isBlackApp(feature.App.PackageName):
			return true
		default:
			return false
		}
	}
}

func isBlackIP(ip string) bool {
	// read db

	// if ip in db, return true
	return false
}

func isBlackApp(pkgName string) bool {
	// read db

	// if pkgName in db, return true
	return false
}

func LowHistoryROI(options ...WhiteTableOption) MediaFilter {
	return func(feature *dbstruct.Feature) bool {
		if reserved(feature, options...) {
			return false
		}
		return false
	}
}
