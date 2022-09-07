package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// UserRead 从数据库中获取用户的信息
func UserRead(context *gin.Context) {
	var user UserID
	err := context.ShouldBindUri(&user)
	if err != nil {
		context.Status(400)
		return
	}

	temp := UsersTable{}
	var searchID string
	searchID = context.Param("id")
	fmt.Println(searchID)

	//tips：first为查询，可以返回查询错误，Find同样为查询，但不能返回错误
	res := Db.Table("users").First(&temp, "user_id = ?", searchID)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(401)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         temp.UserID,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}
