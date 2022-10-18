package views

import (
	"errors"
	"fmt"
	"net/http"
	. "webService_Refactoring/modules"
	"webService_Refactoring/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//TODO 优化1：连接数据库代码冗余... 已解决
//     优化2：命名...

// UpdateUser 更新用户信息，put为上传，patch为修改
func UpdateUser(context *gin.Context) {
	var userID UserID
	var updateUser UpdateUsers
	if err := context.ShouldBindUri(&userID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	if err := context.ShouldBind(&updateUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	temp := UsersTable{}
	searchID := context.Param("id")
	//put方法决定以form-data进行传递数据
	fn := context.PostForm("first_name")
	ln := context.PostForm("last_name")
	if err := Db.Table("users").First(&temp, "user_id = ?", searchID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"msg": utils.GetErrMsg(http.StatusUnauthorized),
		})
		return
	}
	temp.UserFirstName = fn
	temp.UserLastName = ln
	if err := Db.Table("users").Where("user_id = ?", searchID).Updates(&temp).Error; err != nil {
		context.JSON(utils.ErrorUserUpdatedFail, gin.H{
			"msg": utils.GetErrMsg(utils.ErrorUserUpdatedFail),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":        utils.GetErrMsg(http.StatusOK),
		"id":         temp.UserID,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}

// UpdateUserPartial 更新用户部分信息
func UpdateUserPartial(context *gin.Context) {
	var userID UserID
	var updateUser UpdateUsers
	if err := context.ShouldBindUri(&userID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	if err := context.ShouldBind(&updateUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	temp := UsersTable{}
	searchID := context.Param("id")
	newfn := context.PostForm("first_name")
	newln := context.PostForm("last_name")
	if err := Db.Table("users").First(&temp, "user_id = ?", searchID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"msg": utils.GetErrMsg(http.StatusUnauthorized),
		})
		return
	}
	temp.UserFirstName = newfn
	temp.UserLastName = newln
	if err := Db.Table("users").Where("user_id = ?", searchID).Updates(&temp).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":        utils.GetErrMsg(http.StatusOK),
		"id":         temp.UserID,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}
