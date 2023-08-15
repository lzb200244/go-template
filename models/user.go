package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
Created by 斑斑砖 on 2023/8/13.
Description：
	用户模型
*/

const (
	PassWordCost = 12 //密码加密难度
)

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"not null;index;comment:用户名称;"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Email    string `json:"email" gorm:"not null;unique;default:'';comment:邮箱"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

func (user *User) SetPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
