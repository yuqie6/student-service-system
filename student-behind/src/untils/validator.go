package untils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) error {
	pattern := `^[0-9a-zA-Z_.-]+@([0-9a-zA-Z-]+\.)+[a-zA-Z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(email) {
		return fmt.Errorf("邮箱格式不正确")
	}
	return nil
}

// ValidatePhoneNumber 验证手机号码格式
func ValidatePhoneNumber(phone string) error {
	if len(phone) != 11 {
		return fmt.Errorf("手机号码必须是11位数字")
	}
	if phone[0] != '1' {
		return fmt.Errorf("手机号码必须以1开头")
	}
	return nil
}

// ValidateAge 验证年龄合理性
func ValidateAge(age int) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("年龄必须在0-150岁之间")
	}
	return nil
}

// ValidateGender 验证性别输入
func ValidateGender(gender string) error {
	gender = strings.ToLower(strings.TrimSpace(gender))
	if gender != "男" && gender != "女" {
		return fmt.Errorf("性别只能是\"男\"或\"女\"")
	}
	return nil
}
