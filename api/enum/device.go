package enum

type Platform int32

const (
	Platform_Android = 1
	Platform_Ios     = 2
)

var Platform_value = map[int32]string{
	Platform_Android: "android",
	Platform_Ios:     "ios",
}

func (x Platform) String() string {
	return Platform_value[int32(x)]
}

const (
	DeviceTypeMobileTablet = "mobiletablet"
	DeviceTypeComputer     = "computer"
	DeviceTypeTV           = "tv"
	DeviceTypePhone        = "phone"
	DeviceTypeTablet       = "tablet"
)
