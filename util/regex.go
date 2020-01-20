package util

import (
	"regexp"
)

/**
 * 验证Email
 * @param email email地址，格式：zhangsan@sina.com，zhangsan@xxx.com.cn，xxx代表邮件服务商
 * @return 验证成功返回true，验证失败返回false
 */
func IsEmail(email string) (b bool) {
	b, _ = regexp.MatchString(`\\w+@\\w+\\.[a-z]+(\\.[a-z]+)?`, email)
	return
}

/**
 * 验证身份证号码
 * @param idCard 居民身份证号码18位，第一位不能为0，最后一位可能是数字或字母，中间16位为数字 \d同[0-9]
 * @return 验证成功返回true，验证失败返回false
 */
func IsIdCard(idCard string) (b bool) {
	b, _ = regexp.MatchString(`[1-9]\\d{16}[a-zA-Z0-9]{1}`, idCard)
	return
}

/**
 * 验证手机号码（支持国际格式，+86135xxxx...（中国内地），+00852137xxxx...（中国香港））
 * @param mobile 移动、联通、电信运营商的号码段
 *<p>移动的号段：134(0-8)、135、136、137、138、139、147（预计用于TD上网卡）
 *、150、151、152、157（TD专用）、158、159、187（未启用）、188（TD专用）</p>
 *<p>联通的号段：130、131、132、155、156（世界风专用）、185（未启用）、186（3g）</p>
 *<p>电信的号段：133、153、180（未启用）、189</p>
 * @return 验证成功返回true，验证失败返回false
 */
func IsMobile(mobile string) (b bool) {
	b, _ = regexp.MatchString(`13[0-9], 14[5,6,7,8,9], 15[0-3, 5-9], 16[2,5,6,7], 17[0-8], 18[0-9], 19[0-3, 5-9]`, mobile)

	//b, _ = regexp.MatchString(`(\\+\\d+)?1[345789]\\d{9}$`, mobile)
	return
}

/**
 * 验证固定电话号码
 * @param phone 电话号码，格式：国家（地区）电话代码 + 区号（城市代码） + 电话号码，如：+8602085588447
 * <p><b>国家（地区） 代码 ：</b>标识电话号码的国家（地区）的标准国家（地区）代码。它包含从 0 到 9 的一位或多位数字，
 *  数字之后是空格分隔的国家（地区）代码。</p>
 * <p><b>区号（城市代码）：</b>这可能包含一个或多个从 0 到 9 的数字，地区或城市代码放在圆括号——
 * 对不使用地区或城市代码的国家（地区），则省略该组件。</p>
 * <p><b>电话号码：</b>这包含从 0 到 9 的一个或多个数字 </p>
 * @return 验证成功返回true，验证失败返回false
 */
func IsPhone(phone string) (b bool) {
	b, _ = regexp.MatchString(`(\\+\\d+)?(\\d{3,4}\\-?)?\\d{7,8}$`, phone)
	return
}

/**
 * 验证整数（正整数和负整数）
 * @param digit 一位或多位0-9之间的整数
 * @return 验证成功返回true，验证失败返回false
 */
func IsDigit(digit string) (b bool) {
	b, _ = regexp.MatchString(`\\-?[1-9]\\d+`, digit)
	return
}

/**
 * 验证整数和浮点数（正负整数和正负浮点数）
 * @param decimals 一位或多位0-9之间的浮点数，如：1.23，233.30
 * @return 验证成功返回true，验证失败返回false
 */
func IsDecimals(decimals string) (b bool) {
	b, _ = regexp.MatchString(`\\-?[1-9]\\d+(\\.\\d+)?`, decimals)
	return
}

/**
 * 验证空白字符
 * @param blankSpace 空白字符，包括：空格、\t、\n、\r、\f、\x0B
 * @return 验证成功返回true，验证失败返回false
 */
func IsBlankSpace(blankSpace string) (b bool) {
	b, _ = regexp.MatchString(`\\s+`, blankSpace)
	return
}

/**
 * 验证中文
 * @param chinese 中文字符
 * @return 验证成功返回true，验证失败返回false
 */
func IsChinese(chinese string) (b bool) {
	b, _ = regexp.MatchString(`^[\u4E00-\u9FA5]+$`, chinese)
	return
}

/**
 * 验证日期（年月日）
 * @param birthday 日期，格式：1992-09-03，或1992.09.03
 * @return 验证成功返回true，验证失败返回false
 */
func checkBirthday(birthday string) (b bool) {
	b, _ = regexp.MatchString(`[1-9]{4}([-./])\\d{1,2}\\1\\d{1,2}`, birthday)
	return
}

/**
 * 验证URL地址
 * @param url 格式：http://blog.csdn.net:80/xyang81/article/details/7705960? 或 http://www.csdn.net:80
 * @return 验证成功返回true，验证失败返回false
 */
func checkURL(url string) (b bool) {
	b, _ = regexp.MatchString(`(https?://(w{3}\\.)?)?\\w+\\.\\w+(\\.[a-zA-Z]+)*(:\\d{1,5})?(/\\w*)*(\\??(.+=.*)?(&.+=.*)?)?`, url)
	return
}

/**
 * 匹配中国邮政编码
 * @param postcode 邮政编码
 * @return 验证成功返回true，验证失败返回false
 */
func IsPostcode(postcode string) (b bool) {
	b, _ = regexp.MatchString(`[1-9]\\d{5}`, postcode)
	return
}

/**
 * 匹配IP地址(简单匹配，格式，如：192.168.1.1，127.0.0.1，没有匹配IP段的大小)
 * @param latitude IPv4标准地址
 * @return 验证成功返回true，验证失败返回false
 */
func IsIP(ip string) (b bool) {
	b, _ = regexp.MatchString(`^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$`, ip)
	return
}
