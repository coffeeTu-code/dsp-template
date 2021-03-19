package helper

import (
	"bytes"
	"strconv"
	"strings"
)

func GetSdkVersionPrefix(sdk string) string {
	if sdk == "" {
		return ""
	}
	elems := strings.Split(sdk, "_")
	if len(elems) > 1 {
		return elems[0]
	}
	return ""

}

// GetStandardVersion :获取标准版本号
// ver: 获得版本号，如1.1.1、1.1.1_unity、1.1.1+unity等版本号
// 返回值：版本:1.1.1
func GetStandardVersion(ver string) (resVer string) {
	var ok bool
	resVer = ver
	if resVer, ok = SplitVersionAndGetNum(ver, "-"); ok {
		return
	}
	if resVer, ok = SplitVersionAndGetNum(ver, "_"); ok {
		return
	}
	if resVer, ok = SplitVersionAndGetNum(ver, "+"); ok {
		return
	}
	if resVer, ok = SplitVersionAndGetNum(ver, " "); ok {
		return
	}
	return resVer
}

// SplitVersionAndGetNum :用特殊符号分割版本号，并提取数字版本号
// ver: 版本号，如1.1.1、1.1.1_unity、1.1.1+unity等版本号
// subStr:分隔符，如_、-、+
// 返回值：版本:1.1.1、是否split获得
func SplitVersionAndGetNum(ver, subStr string) (string, bool) {
	if strings.Index(ver, subStr) != -1 {
		vers := strings.Split(ver, subStr)
		if len(vers) > 0 && len(vers) <= 2 {
			if len(GetDigitStrFromStr(vers[0])) > 0 {
				return vers[0], true
			}
			if len(GetDigitStrFromStr(vers[1])) > 0 {
				return vers[1], true
			}
		}
	}
	return ver, false
}

//从前到后找出字符串中的数字并返回字符串组成的数字
func GetDigitStrFromStr(str string) string {
	var buffer bytes.Buffer

	for _, s := range str {
		//字符'0'的ascii 48, '9'的ascii 57
		if s >= 48 && s <= 57 {
			buffer.WriteString(string(s))
		}

	}
	return buffer.String()
}

//计算方式：取osv前4位，不足补0，每位按2个位置计算。13.5.3 -> 13.05.03.00 -> 13050300
func GetSdkVersionNum(sdk string) (sdkNum int) {
	return GetVersionCode(GetStandardVersion(sdk))
}

func GetAsSdkVersionNum(sdk string) (sdkNum int) {
	if sdk == "" {
		return 0
	}

	elems := strings.Split(sdk, "_")
	if len(elems) > 1 {
		return GetAsVersionCode(elems[1])
	}
	return 0
}

//计算方式：取osv前4位，不足补0，每位按2个位置计算。13.5.3 -> 13.05.03.00 -> 13050300
func GetVersionCode(version string) (code int) {
	items := strings.Split(version, ".")
	var base int
	var err error
	for i := 0; i < 4; i++ {
		base = 0
		if i < len(items) {
			base, err = strconv.Atoi(items[i])
			if err != nil {
				return 0
			}
		}
		code = code*100 + base
	}
	return code
}

func GetAsVersionCode(version string) (code int) {
	items := strings.Split(version, ".")
	var base int
	var err error
	for i := 0; i < 3; i++ {
		base = 0
		if i < len(items) {
			base, err = strconv.Atoi(items[i])
			if err != nil {
				return 0
			}
		}
		code = code*1000 + base
	}
	return code
}
