package untils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateEmail 验证邮箱格式
// 参数：
// - email：待验证的邮箱地址
// 返回值：
// - error：如果邮箱格式不正确，返回错误信息；否则返回nil
func ValidateEmail(email string) error {
	pattern := `^[0-9a-zA-Z_.-]+@([0-9a-zA-Z-]+\.)+[a-zA-Z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(email) {
		return fmt.Errorf("邮箱格式不正确")
	}
	return nil
}

// ValidatePhoneNumber 验证手机号码格式
// 参数：
// - phone：待验证的手机号码
// 返回值：
// - error：如果手机号码格式不正确，返回错误信息；否则返回nil
func ValidatePhoneNumber(phone int) error {
	// 转换为字符串进行验证
	phoneStr := fmt.Sprintf("%d", phone)
	if len(phoneStr) != 11 {
		return fmt.Errorf("手机号码必须是11位数字")
	}
	if phoneStr[0] != '1' {
		return fmt.Errorf("手机号码必须以1开头")
	}
	return nil
}

// ValidateAge 验证年龄合理性
// 参数：
// - age：待验证的年龄
// 返回值：
// - error：如果年龄不在合理范围内，返回错误信息；否则返回nil
func ValidateAge(age int) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("年龄必须在0-150岁之间")
	}
	return nil
}

// ValidateGender 验证性别输入
// 参数：
// - gender：待验证的性别
// 返回值：
// - error：如果性别输入不正确，返回错误信息；否则返回nil
func ValidateGender(gender string) error {
	gender = strings.ToLower(strings.TrimSpace(gender))
	if gender != "男" && gender != "女" {
		return fmt.Errorf("性别只能是\"男\"或\"女\"")
	}
	return nil
}
