package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func initDB() {
	// 修改为 MySQL 数据库连接
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/bolg?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()

	// 定义用户相关路由
	r.POST("/register", registerUser)
	r.POST("/login", loginUser)

	// 文章相关路由
	postGroup := r.Group("/posts")
	{
		// 创建文章（需要认证）
		postGroup.POST("", authMiddleware(), createPost)
		// 获取所有文章
		postGroup.GET("", getPosts)
		// 获取单篇文章
		postGroup.GET("/:id", getPost)
		// 更新文章（需要认证）
		postGroup.PUT("/:id", authMiddleware(), updatePost)
		// 删除文章（需要认证）
		postGroup.DELETE("/:id", authMiddleware(), deletePost)
	}

	// 评论相关路由
	commentGroup := r.Group("/comments")
	{
		// 创建评论（需要认证）
		commentGroup.POST("/:post_id", authMiddleware(), createComment)
		// 获取某篇文章的所有评论
		commentGroup.GET("/:post_id", getCommentsByPost)
	}

	r.Run(":8080")
}
