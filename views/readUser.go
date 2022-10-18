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

// UserRead 从数据库中获取用户的信息
func UserRead(context *gin.Context) {
	var user UserID
	if err := context.ShouldBindUri(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	temp := UsersTable{}
	searchID := context.Param("id")
	fmt.Println(searchID)
	//tips：first为查询，可以返回查询错误，Find同样为查询，但不能返回错误
	if err := Db.Table("users").First(&temp, "user_id = ?", searchID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"msg": utils.GetErrMsg(http.StatusUnauthorized),
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
