package service

import (
	"fmt"
	"strconv"
	"strings"
	"studentservice/models"
	"studentservice/untils"
)

// 定义数据文件的路径
var dataFile = "students.json"

// AddStudent 函数用于添加新的学生信息到学生列表
// 参数：
// - students：指向学生列表的指针
// 返回值：
// - error：可能的错误信息
func AddStudent(students *[]models.Student) error {
	for {
		fmt.Println("=== 添加学生 ===")
		// 提示用户输入学生的ID，输入0可以返回主菜单
		id1, err := untils.GetIntInput("请输入要添加的学生的Id（输入0返回主菜单）：")
		if err != nil {
			fmt.Println("输入的Id错误：", err)
			continue
		}
		if id1 == 0 {
			fmt.Println("已取消添加学生，返回主菜单。")
			return nil
		}
		// 检查ID是否已存在
		var exists bool
		for _, v := range *students {
			if v.Id == id1 {
				fmt.Println("已存在这个学生!")
				exists = true
				break
			}
		}
		if exists {
			// 询问用户是否重新输入ID或返回主菜单
			choice, err := untils.GetStringInput("是否重新输入ID？（y/n）：")
			if err != nil {
				fmt.Println("输入错误：", err)
				continue
			}
			if strings.ToLower(choice) == "n" {
				fmt.Println("已取消添加学生，返回主菜单。")
				return nil
			} else {
				continue
			}
		}
		// 获取其他学生信息
		age1, err := getValidatedIntInput("请输入要添加的学生的年龄:", untils.ValidateAge)
		if err != nil {
			fmt.Println("错误:", err)
			continue
		}

		phoneNumber1, err := getValidatedStringInput("请输入要添加的学生的电话号码:", untils.ValidatePhoneNumber)
		if err != nil {
			fmt.Println("错误:", err)
			continue
		}

		name1, err := untils.GetStringInput("请输入要添加的学生的名字:")
		if err != nil {
			fmt.Println("错误:", err)
			continue
		}

		gender1, err := getValidatedStringInput("请输入要添加的学生的性别(男/女):", untils.ValidateGender)
		if err != nil {
			fmt.Println("错误:", err)
			continue
		}

		email1, err := getValidatedStringInput("请输入要添加的学生的电子邮箱:", untils.ValidateEmail)
		if err != nil {
			fmt.Println("错误:", err)
			continue
		}

		// 创建新的学生实例
		newstudent := models.Student{
			Id:          id1,
			Name:        name1,
			Age:         age1,
			Gender:      gender1,
			Email:       email1,
			PhoneNumber: phoneNumber1,
		}
		// 将新学生添加到列表
		*students = append(*students, newstudent)
		fmt.Println("添加成功!")

		// 保存学生数据到文件
		err = models.SaveStudents(*students, dataFile)
		if err != nil {
			fmt.Println("保存学生数据失败：", err)
		}
		break
	}
	return nil
}

// DeleteStudent 函数用于删除指定ID的学生信息
// 参数：
// - students：指向学生列表的指针
// 返回值：
// - error：可能的错误信息
func DeleteStudent(students *[]models.Student) error {
	fmt.Println("=== 删除学生 ===")
	id1, err := untils.GetIntInput("请输入要删除的学生Id：")
	if err != nil {
		fmt.Println("错误:", err)
		return err
	}
	for i, v := range *students {
		if v.Id == id1 {
			// 找到匹配的学生，删除并退出函数
			*students = append((*students)[:i], (*students)[i+1:]...)
			fmt.Println("删除学生成功!")

			// 保存学生数据到文件
			err = models.SaveStudents(*students, dataFile)
			if err != nil {
				fmt.Println("保存学生数据失败：", err)
			}
			return nil
		}
	}
	// 未找到匹配的学生
	fmt.Println("不存在这个学生!")
	return fmt.Errorf("学生不存在")
}

// ViewStudent 函数用于查看指定ID的学生信息
// 参数：
// - students：指向学生列表的指针
// 返回值：
// - error：可能的错误信息
func ViewStudent(students *[]models.Student) error {
	fmt.Println("=== 查找学生 ===")
	id, err := untils.GetIntInput("请输入要查找的学生Id：")
	if err != nil {
		fmt.Println("错误:", err)
		return err
	}
	var check bool
	for _, v := range *students {
		if v.Id == id {
			check = true
			fmt.Printf("ID: %d\n名字: %s\n年龄: %d\n性别: %s\n邮箱: %s\n电话: %s\n",
				v.Id, v.Name, v.Age, v.Gender, v.Email, v.PhoneNumber)
		}
	}
	if !check {
		fmt.Println("不存在这个学生!")
		return fmt.Errorf("学生不存在")
	}
	fmt.Println("查找完毕！")
	return nil
}

// EditStudent 函数用于修改指定ID的学生信息
// 参数：
// - students：指向学生列表的指针
// 返回值：
// - error：可能的错误信息
func EditStudent(students *[]models.Student) error {
	fmt.Println("=== 编辑学生 ===")
	id, err := untils.GetIntInput("请输入要修改的学生Id：")
	if err != nil {
		fmt.Println("错误:", err)
		return err
	}
	var check bool
	var tag int
	for i, v := range *students {
		if v.Id == id {
			check = true
			tag = i
			break
		}
	}
	if !check {
		fmt.Println("不存在这个学生!")
		return fmt.Errorf("学生不存在")
	}
	for {
		// 显示修改选项菜单
		fmt.Println("请选择要进行的操作：")
		fmt.Println("1.修改学生姓名")
		fmt.Println("2.修改学生电话号码")
		fmt.Println("3.修改学生邮箱")
		fmt.Println("4.修改学生性别")
		fmt.Println("5.修改学生年龄")
		fmt.Println("6.完成修改")
		var choice int
		for {
			// 获取用户的选择
			input, err := untils.GetStringInput("请输入你的选择（1-6）：")
			if err != nil {
				fmt.Println("输入无效，请输入 1 到 6 之间的数字。")
				continue
			}

			// 转换为整数
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("输入无效，请输入 1 到 6 之间的数字。")
				continue
			}

			if num < 1 || num > 6 {
				fmt.Println("输入无效，请输入 1 到 6 之间的数字。")
				continue
			}

			choice = num
			break
		}
		// 根据用户的选择执行相应的修改操作
		switch choice {
		case 1:
			// 修改学生姓名
			name, err := untils.GetStringInput("请输入修改后的名字:")
			if err != nil {
				fmt.Printf("操作失败: %v\n", err)
				continue
			} else {
				(*students)[tag].Name = name
			}
		case 2:
			// 修改学生电话号码
			phoneNumber, err := getValidatedStringInput("请输入修改后的电话号码:", untils.ValidatePhoneNumber)
			if err != nil {
				fmt.Printf("操作失败: %v\n", err)
				continue
			}
			(*students)[tag].PhoneNumber = phoneNumber
		case 3:
			// 修改学生邮箱
			email, err := getValidatedStringInput("请输入修改后的邮箱:", untils.ValidateEmail)
			if err != nil {
				fmt.Printf("操作失败: %v\n", err)
				continue
			}
			(*students)[tag].Email = email
		case 4:
			// 修改学生性别
			gender, err := getValidatedStringInput("请输入修改后的性别(男/女):", untils.ValidateGender)
			if err != nil {
				fmt.Printf("操作失败: %v\n", err)
				continue
			}
			(*students)[tag].Gender = gender
		case 5:
			// 修改学生年龄
			age, err := getValidatedIntInput("请输入修改后的年龄:", untils.ValidateAge)
			if err != nil {
				fmt.Printf("操作失败: %v\n", err)
				continue
			}
			(*students)[tag].Age = age
		case 6:
			// 完成修改
			fmt.Println("修改完成")

			// 保存学生数据到文件
			err := models.SaveStudents(*students, dataFile)
			if err != nil {
				fmt.Println("保存学生数据失败：", err)
			}
			return nil
		}
	}
}

// getValidatedIntInput 获取并验证整数输入
// 参数：
// - prompt：提示信息
// - validateFunc：验证函数
// 返回值：
// - int：用户输入的整数
// - error：可能的错误信息
func getValidatedIntInput(prompt string, validateFunc func(int) error) (int, error) {
	for {
		input, err := untils.GetIntInput(prompt)
		if err != nil {
			fmt.Println("输入错误：", err)
			continue
		}
		if err := validateFunc(input); err != nil {
			fmt.Println("验证错误：", err)
			continue
		}
		return input, nil
	}
}

// getValidatedStringInput 获取并验证字符串输入
// 参数：
// - prompt：提示信息
// - validateFunc：验证函数
// 返回值：
// - string：用户输入的字符串
// - error：可能的错误信息
func getValidatedStringInput(prompt string, validateFunc func(string) error) (string, error) {
	for {
		input, err := untils.GetStringInput(prompt)
		if err != nil {
			fmt.Println("输入错误：", err)
			continue
		}
		if err := validateFunc(input); err != nil {
			fmt.Println("验证错误：", err)
			continue
		}
		return input, nil
	}
}

func GetAllStudents() ([]models.Student, error) {
	return models.LoadStudents(dataFile)
}

func GetStudentById(id int) (*models.Student, error) {
	students, err := models.LoadStudents(dataFile)
	if err != nil {
		return nil, err
	}

	for _, student := range students {
		if student.Id == id {
			return &student, nil
		}
	}
	return nil, fmt.Errorf("学生不存在")
}

func AddNewStudent(student *models.Student) error {
	if err := validateStudent(student); err != nil {
		return err
	}

	students, err := models.LoadStudents(dataFile)
	if err != nil {
		return err
	}

	// 检查ID是否已存在
	for _, s := range students {
		if s.Id == student.Id {
			return fmt.Errorf("学生ID已存在")
		}
	}

	students = append(students, *student)
	return models.SaveStudents(students, dataFile)
}

func UpdateStudent(student *models.Student) error {
	if err := validateStudent(student); err != nil {
		return err
	}

	students, err := models.LoadStudents(dataFile)
	if err != nil {
		return err
	}

	for i, s := range students {
		if s.Id == student.Id {
			students[i] = *student
			return models.SaveStudents(students, dataFile)
		}
	}
	return fmt.Errorf("学生不存在")
}

func DeleteStudentById(id int) error {
	students, err := models.LoadStudents(dataFile)
	if err != nil {
		return err
	}

	for i, student := range students {
		if student.Id == id {
			students = append(students[:i], students[i+1:]...)
			return models.SaveStudents(students, dataFile)
		}
	}
	return fmt.Errorf("学生不存在")
}

func validateStudent(student *models.Student) error {
	if err := untils.ValidateAge(student.Age); err != nil {
		return err
	}
	if err := untils.ValidatePhoneNumber(student.PhoneNumber); err != nil {
		return err
	}
	if err := untils.ValidateEmail(student.Email); err != nil {
		return err
	}
	if err := untils.ValidateGender(student.Gender); err != nil {
		return err
	}
	return nil
}
