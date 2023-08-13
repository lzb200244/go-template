package models

import "gorm.io/gorm"

/*
Created by 斑斑砖 on 2023/8/13.
Description：
	用户模型
*/

type User struct {
	gorm.Model
	UserName string `json:"name" gorm:"not null;index;comment:用户名称;"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Email    string `json:"email" gorm:"not null;unique;default:'';comment:邮箱"`
}
