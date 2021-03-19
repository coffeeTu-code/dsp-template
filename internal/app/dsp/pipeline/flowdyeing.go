package pipeline

import (
	"net"
	"strings"

	"dsp-template/api/dbstruct"

	"dsp-template/api/base"
	dsp_config "dsp-template/internal/app/dsp/dsp-config"
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
	abtesting "dsp-template/pkg2/helper-abtesting"
	helper_crypto "dsp-template/pkg2/helper-crypto"
)

func NewFlowDyeingPipeline() *FlowDyeingPipeline {
	return &FlowDyeingPipeline{
		funcBaseDyeing: []Dyeing{
			DyeingRequestBase,
			DyeingABTestKV,
		},
	}
}

type FlowDyeingPipeline struct {
	funcBaseDyeing []Dyeing
}

func (pipe *FlowDyeingPipeline) Description() string {
	return "FlowDyeing"
}

func (pipe *FlowDyeingPipeline) Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelErr error) {
	runDyeing(ctx, pipe.funcBaseDyeing)
	return "", nil
}

func runDyeing(ctx *dsp_context.DspContext, list []Dyeing) {
	for _, dyeing := range list {
		dyeing(ctx)
	}
}

type Dyeing func(ctx *dsp_context.DspContext)

func DyeingABTestKV(ctx *dsp_context.DspContext) {
	kvs := abtesting.APP().GetAbTesting(&abtesting.FlowInfo{
		Region:      ctx.MetricsCtx.Base.Region,
		Adx:         ctx.Feature.Exchange,
		AdType:      ctx.Feature.AdType,
		Device:      ctx.Feature.Platform,
		Country:     ctx.Feature.CountryCode,
		PackageName: ctx.Feature.App.PackageName, //流量测包名
		HashKey:     hashkey(ctx.Feature),
	})
	for k, v := range kvs {
		ctx.RequestBase.Abtest[k] = v
	}
}

func hashkey(feature *base.Feature) string {
	var hashKey string
	if hashKey = uniqDeviceId(feature.DeviceIds); hashKey != "" {
		return hashKey
	}
	if hashKey = helper_crypto.Md5(ipua(feature.Device)); hashKey != "" {
		return hashKey
	}
	return ""
}

//获得可用设备ID
func uniqDeviceId(ids *dbstruct.FDeviceIds) string {
	deviceIdList := []string{
		ids.Idfa,
		ids.GoogleAdId,
		ids.GoogleAdIdMD5,
		ids.Imei,
		ids.ImeiMD5,
		ids.AndroidId,
		ids.AndroidIdMD5,
		ids.OAId,
		ids.OAIdMD5,
	}
	for i := range deviceIdList {
		if validDeviceId(deviceIdList[i]) {
			return deviceIdList[i]
		}
	}
	return ""
}

// 判断设备id是否合法
// 非法设备ID 如 deviceId == "" || deviceId == "-" || deviceId == "00000000-0000-0000-0000-000000000000"
// return bool 合法返回true, 非法返回false
func validDeviceId(deviceId string) bool {
	tmpDeviceId := strings.TrimSpace(deviceId)
	if tmpDeviceId == "" || tmpDeviceId == "null" || tmpDeviceId == "gaid" || tmpDeviceId == "idfa" {
		return false
	}
	//判断是否为无效设备id，返回fasle
	// read config
	// if tmpDeviceId in map, return false
	for i := 0; i < len(tmpDeviceId); i++ {
		if tmpDeviceId[i] != '0' && tmpDeviceId[i] != '-' && tmpDeviceId[i] != ' ' {
			return true
		}
	}
	return false
}

func ipua(device *base.FDevice) string {
	if len(device.UserAgent) == 0 {
		return ""
	}
	if !validIP(device.IP) {
		return ""
	}
	return strings.Join([]string{device.IP, device.UserAgent}, "#")
}

// 判断IP地址是否为合法IP，排除内网ip
func validIP(s string) bool {
	ip := net.ParseIP(strings.TrimSpace(s))
	if !ip.IsGlobalUnicast() {
		return false
	}

	if ip4 := ip.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		case ip4[0] >= 224:
			return false
		default:
			return true
		}
	}
	if ip16 := ip.To16(); ip16 != nil {
		return ip16[0] != 0xfd || ip16[1]&0xc0 != 0x00
	}
	return false
}

func DyeingRequestBase(ctx *dsp_context.DspContext) {
	ctx.RequestBase.RequestId = ctx.Request.ID
	ctx.RequestBase.Debug = debug(ctx)
	ctx.RequestBase.Openlog = openlog(ctx)
}

func debug(ctx *dsp_context.DspContext) bool {
	var _debug bool
	return _debug
}

func openlog(ctx *dsp_context.DspContext) bool {
	var _openlog bool
	force := dsp_config.GetDspConfig().DebugConfig.BidForce
	if force != nil && len(force.BidForceDevice) > 0 {
		ctx.BidForce = bidforce(ctx, force)
		_openlog = ctx.BidForce.IsForce()
	}
	return _openlog
}

func bidforce(ctx *dsp_context.DspContext, force *dsp_config.BidForce) *dsp_config.BidForceDevice {
	if ctx.Feature == nil || ctx.Feature.DeviceIds == nil || len(force.BidForceDevice) == 0 {
		return nil
	}
	deviceids := ctx.Feature.DeviceIds
	devicemap := map[string]bool{
		deviceids.GoogleAdId:     true, //1
		deviceids.Idfa:           true, //2
		deviceids.GoogleAdIdMD5:  true, //3
		deviceids.GoogleAdIdSHA1: true, //4
		deviceids.Imei:           true, //5
		deviceids.ImeiMD5:        true, //6
		deviceids.ImeiSHA1:       true, //7
		deviceids.AndroidId:      true, //8
		deviceids.AndroidIdMD5:   true, //9
		deviceids.AndroidIdSHA1:  true, //10
		deviceids.OAId:           true, //11
		deviceids.OAIdMD5:        true, //12
	}
	for user, forcedevice := range force.BidForceDevice {
		if forcedevice == nil {
			continue
		}
		for _, id := range forcedevice.DeviceId {
			if devicemap[id] {
				forcedevice.User = user
				return forcedevice
			}
		}
	}
	return nil
}
