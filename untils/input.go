package untils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// GetIntInput 函数用于获取用户输入的整数
// 参数：
// - prompt：提示用户输入的信息
// 返回值：
// - int：用户输入的整数
// - error：可能的错误信息
func GetIntInput(prompt string) (int, error) {
	if prompt != "" {
		fmt.Println(prompt)
	}
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, fmt.Errorf("输入不能为空:%w", err)
	}
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("转换失败:%w", err)
	}
	if number < 0 {
		return 0, fmt.Errorf("请输入一个正确的数字")
	}
	return number, nil
}

// GetStringInput 函数用于获取用户输入的字符串
// 参数：
// - prompt：提示用户输入的信息
// 返回值：
// - string：用户输入的字符串
// - error：可能的错误信息
func GetStringInput(prompt string) (string, error) {
	if prompt != "" {
		fmt.Println(prompt)
	}
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	if input == "" {
		return "", fmt.Errorf("输入不能为空")
	}
	return input, nil
}
