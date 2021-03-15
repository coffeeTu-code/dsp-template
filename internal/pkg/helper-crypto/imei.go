package helper_crypto

import "strconv"

/*

IMEI本该最理想的设备ID，具备唯一性，恢复出厂设置不会变化（真正的设备相关）。
然而，获取IMEI需要 READ_PHONE_STATE 权限，估计大家也知道这个权限有多麻烦了。
尤其是Android 6.0以后, 这类权限要动态申请，很多用户可能会选择拒绝授权。
我们看到，有的APP不授权这个权限就无法使用， 这可能会降低用户对APP的好感度。


而且，Android 10.0 将彻底禁止第三方应用获取设备的IMEI, 即使申请了 READ_PHONE_STATE 权限。
所以，如果是新APP，不建议用IMEI作为设备标识；
如果已经用IMEI作为标识，要赶紧做兼容工作了，尤其是做新设备标识和IMEI的映射。

*/

/*

IMEI校验码算法：

(1).将偶数位数字分别乘以2，分别计算个位数和十位数之和
(2).将奇数位数字相加，再加上上一步算得的值
(3).如果得出的数个位是0则校验位为0，否则为10减去个位数
如：35 89 01 80 69 72 41 偶数位乘以2得到5*2=10 9*2=18 1*2=02 0*2=00 9*2=18 2*2=04 1*2=02,计算奇数位数字之和和偶数位个位十位之和，
  得到 3+(1+0)+8+(1+8)+0+(0+2)+8+(0+0)+6+(1+8)+7+(0+4)+4+(0+2)=63 => 校验位 10-3 = 7，则最后位为7，所以完整是358901806972417

*/

func calculateCheckDigit(prefix string) int {
	var total, sum1, sum2 int
	n := len(prefix)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(string(prefix[i]))
		// 奇数
		if i % 2 == 0 {
			sum1 += num
		} else { // 偶数
			tmp := num * 2
			if tmp < 10 {
				sum2 += tmp
			} else {
				sum2 = sum2 + tmp + 1 - 10
			}
		}
	}
	total = sum1 + sum2
	if (total % 10 == 0) {
		return  0
	} else {
		return  10 - (total % 10)
	}
}
