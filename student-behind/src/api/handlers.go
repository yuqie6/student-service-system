package api

import (
	"net/http"
	"strconv"
	"studentservice/models"
	"studentservice/service"

	"github.com/gin-gonic/gin"
)

// CreateStudent 处理创建学生请求
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.AddNewStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// GetStudent 处理获取单个学生信息请求
func GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	student, err := service.GetStudentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该学生"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// UpdateStudent 处理更新学生信息请求
func UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.Id = id
	if err := service.UpdateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudent 处理删除学生请求
func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := service.DeleteStudentById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// GetAllStudents 处理获取所有学生信息请求
func GetAllStudents(c *gin.Context) {
	students, err := service.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}
