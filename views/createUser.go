package views

import (
	"net/http"
	. "webService_Refactoring/modules"
	"webService_Refactoring/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// UserCreate 生成用户数据并存储到数据库中
func UserCreate(context *gin.Context) {
	//注册命名规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("usernamerule", UsernameRule)
	}
	var registerForm RegisterForm
	if err := context.ShouldBind(&registerForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	tokenKey := registerForm.Username + "&" + registerForm.Password
	token := MD5(tokenKey)
	userUUID := CreateUUID()

	DbCreateUser := UsersTable{
		UserID:        userUUID,
		UserName:      registerForm.Username,
		UserToken:     token,
		UserPassword:  MD5(registerForm.Password), //暂定md5密文存储密码，存在一些问题，与导师商量后再定
		UserFirstName: registerForm.Firstname,
		UserLastName:  registerForm.Lastname,
		UserEmail:     registerForm.Email,
	}
	if err := Db.Table("users").Create(&DbCreateUser).Error; err != nil {
		context.JSON(utils.ErrorUserCreatedFail, gin.H{
			"msg": utils.GetErrMsg(utils.ErrorUserCreatedFail),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"msg":        utils.GetErrMsg(http.StatusOK),
		"id":         userUUID.String(),
		"username":   registerForm.Username,
		"password":   registerForm.Password,
		"first_name": registerForm.Firstname,
		"last_name":  registerForm.Lastname,
		"email":      registerForm.Email,
		"auth_token": token,
	})

}
