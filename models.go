package main

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username     string `json:"username"`
	PasswordHash string `json:"-"` // 不会在 JSON 响应中返回
	Email        string `json:"email"`
}

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
}

// SetPassword 设置用户密码
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

// CheckPassword 验证用户密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}