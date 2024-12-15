package models

// Student 结构体表示学生的信息
type Student struct {
	Id          int    `json:"id"`           // 学生的唯一标识符
	Name        string `json:"name"`         // 学生的姓名
	Age         int    `json:"age"`          // 学生的年龄
	Gender      string `json:"gender"`       // 学生的性别
	Email       string `json:"email"`        // 学生的电子邮箱
	PhoneNumber int    `json:"phone_number"` // 学生的电话号码
}
