package main

import (
	"fmt"
	"studentservice/models"
	"studentservice/service"
	"studentservice/untils"
)

// main 函数是程序的入口，显示菜单并处理用户的选择
func main() {
	// 定义数据文件的路径
	dataFile := "students.json"

	// 从文件加载学生数据
	students, err := models.LoadStudents(dataFile)
	if err != nil {
		fmt.Println("加载学生数据失败：", err)
		return
	}

	fmt.Println("=== 欢迎来到学生管理系统 ===")
	for {
		fmt.Println("请选择要进行的操作：")
		fmt.Println("1. 添加学生")
		fmt.Println("2. 查看学生")
		fmt.Println("3. 修改学生")
		fmt.Println("4. 删除学生")
		fmt.Println("5. 退出系统")
		var choice int
		choice, err := untils.GetIntInput("请输入你的选择（1-5）：")
		if err != nil {
			fmt.Println("输入无效，请输入 1 到 5 之间的数字。")
			continue
		}
		switch choice {
		case 1:
			// 添加学生
			service.AddStudent(&students)
		case 2:
			// 查看学生
			service.ViewStudent(&students)
		case 3:
			// 修改学生
			service.EditStudent(&students)
		case 4:
			// 删除学生
			service.DeleteStudent(&students)
		case 5:
			// 程序退出前保存学生数据到文件
			err := models.SaveStudents(students, dataFile)
			if err != nil {
				fmt.Println("保存学生数据失败：", err)
			} else {
				fmt.Println("学生数据已保存。")
			}
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}