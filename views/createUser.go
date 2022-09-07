package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	. "webService_Refactoring/modules"
)

// UserCreate 生成用户数据并存储到数据库中
func UserCreate(context *gin.Context) {
	//注册命名规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("usernamerule", UsernameRule)
	}
	var registerForm RegisterForm
	err := context.ShouldBind(&registerForm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var tokenKey string
	tokenKey = registerForm.Username + "&" + registerForm.Password
	token := MD5(tokenKey)

	var userUUID uuid.UUID
	userUUID = CreateUUID()

	context.JSON(http.StatusOK, gin.H{
		"id":         userUUID.String(),
		"username":   registerForm.Username,
		"password":   registerForm.Password,
		"first_name": registerForm.Firstname,
		"last_name":  registerForm.Lastname,
		"email":      registerForm.Email,
		"auth_token": token,
	})

	DbCreateUser := UsersTable{
		UserID:        userUUID,
		UserName:      registerForm.Username,
		UserToken:     token,
		UserPassword:  MD5(registerForm.Password), //暂定md5密文存储密码，存在一些问题，与导师商量后再定
		UserFirstName: registerForm.Firstname,
		UserLastName:  registerForm.Lastname,
		UserEmail:     registerForm.Email,
	}
	fmt.Println(Db.Table("users").Create(&DbCreateUser).RowsAffected)
}
