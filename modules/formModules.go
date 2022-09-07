package modules

//这里是post方法中定义表单的结构体

import (
	"github.com/go-playground/validator/v10"
)

// LoginForm 登录表单
type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}

// RegisterForm 用户注册表单
type RegisterForm struct {
	Username  string `form:"username" json:"username" binding:"required,min=1,max=150,usernamerule"`
	Password  string `form:"password" json:"password" binding:"required,min=1,max=120"`
	Firstname string `form:"first_name" json:"first_name" binding:"max=150"`
	Lastname  string `form:"last_name" json:"last_name" binding:"max=150"`
	Email     string `form:"email" json:"email" binding:"required,email,max=254"`
}

// UsernameRule 用户名规则
var UsernameRule validator.Func = func(fl validator.FieldLevel) bool {
	s, _ := fl.Field().Interface().(string)
	for _, ch := range s {
		if !((ch >= 48 && ch <= 57) || (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122) ||
			ch == 64 || ch == 43 || ch == 45 || ch == 46 || ch == 95) {
			return false
		}
	}
	return true
}

// UserID 用户ID,采用uuid
type UserID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// UpdateUsers 更新用户
type UpdateUsers struct {
	FirstName string `form:"first_name" binding:"required,max=150"`
	LastName  string `form:"last_name" binding:"required,max=150"`
}

// ProjectID 项目ID
type ProjectID struct {
	Pid string `form:"pid" binding:"required"`
}
