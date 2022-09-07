package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

//TODO 优化1：连接数据库代码冗余... 已解决
//     优化2：命名...

// UpdateUser 更新用户信息，put为上传，patch为修改
func UpdateUser(context *gin.Context) {
	var userID UserID
	err := context.ShouldBindUri(&userID)
	if err != nil {
		context.Status(400)
		return
	}
	var updateUser UpdateUsers
	errs := context.ShouldBind(&updateUser)
	if errs != nil {
		context.Status(400)
		return
	}

	temp := UsersTable{}
	var searchID string
	searchID = context.Param("id")
	fmt.Println(searchID)
	//put方法决定以form-data进行传递数据

	fn := context.PostForm("first_name")
	ln := context.PostForm("last_name")

	res := Db.Table("users").First(&temp, "user_id = ?", searchID)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(401)
		return
	}
	temp.UserFirstName = fn
	temp.UserLastName = ln

	er := Db.Table("users").Where("user_id = ?", searchID).Updates(&temp).Error
	if er != nil {
		fmt.Println(er.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":         temp.UserID,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}

// UpdateUserPartial 更新用户部分信息
func UpdateUserPartial(context *gin.Context) {
	var userID UserID
	err := context.ShouldBindUri(&userID)
	if err != nil {
		context.Status(400)
		return
	}
	var updateUser UpdateUsers
	errs := context.ShouldBind(&updateUser)
	if errs != nil {
		context.Status(400)
		return
	}

	temp := UsersTable{}
	var searchID string
	searchID = context.Param("id")
	fmt.Println(searchID)

	newfn := context.PostForm("first_name")
	newln := context.PostForm("last_name")

	res := Db.Table("users").First(&temp, "user_id = ?", searchID)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(401)
		return
	}
	temp.UserFirstName = newfn
	temp.UserLastName = newln

	er := Db.Table("users").Where("user_id = ?", searchID).Updates(&temp).Error
	if er != nil {
		fmt.Println(er.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         temp.UserID,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}
