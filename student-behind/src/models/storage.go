package models

import (
	"encoding/json"
	"os"
)

// SaveStudents 将学生列表保存到指定的文件中
// 参数：
// - students：学生列表
// - filename：保存文件的路径
// 返回值：
// - error：可能的错误信息
func SaveStudents(students []Student, filename string) error {
	data, err := json.Marshal(students)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadStudents 从指定的文件中加载学生列表
// 参数：
// - filename：加载文件的路径
// 返回值：
// - []Student：学生列表
// - error：可能的错误信息
func LoadStudents(filename string) ([]Student, error) {
	// 如果文件不存在，返回空的学生列表
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Student{}, nil
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var students []Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}
