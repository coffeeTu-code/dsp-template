package enum

// 枚举含义 : 网络类型 As 体系
//********************* NetWork Type *********************
type NetWorkType_As int32

func (network NetWorkType_As) String() string {
	return NetWork_Type_As_name[int32(network)]
}

const (
	NetWorkType_As_UNKNOWN           = 0
	NetWorkType_As_NET_2G            = 2
	NetWorkType_As_NET_3G            = 3
	NetWorkType_As_NET_4G            = 4
	NetWorkType_As_NET_WIFI          = 9
	NetWorkType_As_RECALL_IF_UNKNOWN = 10
)

// 枚举含义 : 网络类型 Dsp 体系
//********************* NetWork Type *********************
const (
	NetWorkType_Dsp_UNKNOWN      = 0
	NetWorkType_Dsp_ETHERNET     = 1
	NetWorkType_Dsp_WIFI         = 2
	NetWorkType_Dsp_CELL_UNKNOWN = 3
	NetWorkType_Dsp_CELL_2G      = 4
	NetWorkType_Dsp_CELL_3G      = 5
	NetWorkType_Dsp_CELL_4G      = 6
)

var NewWork_Type_to_as_value = map[int32]int32{
	NetWorkType_Dsp_UNKNOWN:      NetWorkType_As_UNKNOWN,
	NetWorkType_Dsp_ETHERNET:     NetWorkType_As_UNKNOWN,
	NetWorkType_Dsp_WIFI:         NetWorkType_As_NET_WIFI,
	NetWorkType_Dsp_CELL_UNKNOWN: NetWorkType_As_RECALL_IF_UNKNOWN,
	NetWorkType_Dsp_CELL_2G:      NetWorkType_As_NET_2G,
	NetWorkType_Dsp_CELL_3G:      NetWorkType_As_NET_3G,
	NetWorkType_Dsp_CELL_4G:      NetWorkType_As_NET_4G,
}

var NetWork_Type_As_name = map[int32]string{
	NetWorkType_As_UNKNOWN:           "unknown",
	NetWorkType_As_NET_WIFI:          "wifi",
	NetWorkType_As_RECALL_IF_UNKNOWN: "unknown",
	NetWorkType_As_NET_2G:            "2g",
	NetWorkType_As_NET_3G:            "3g",
	NetWorkType_As_NET_4G:            "4g",
}

func GetNewWorkTypeDspToAs(dspNetWorkType int32) (asNetWorkType int32) {
	return NewWork_Type_to_as_value[dspNetWorkType]
}
