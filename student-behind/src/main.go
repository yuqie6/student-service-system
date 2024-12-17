package main

import (
	"log"
	"studentservice/api"
	"studentservice/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 使用CORS中间件
	r.Use(middleware.CORSMiddleware())

	// API路由组
	v1 := r.Group("/api")
	{
		// 学生相关路由
		v1.GET("/students", api.GetAllStudents)
		v1.GET("/students/:id", api.GetStudent)
		v1.POST("/students", api.CreateStudent)
		v1.PUT("/students/:id", api.UpdateStudent)
		v1.DELETE("/students/:id", api.DeleteStudent)
	}

	// 打印启动信息
	log.Println("服务器启动在 http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("启动服务器失败：%v", err)
	}
}
