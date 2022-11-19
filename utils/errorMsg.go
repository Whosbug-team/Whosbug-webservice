package utils

const (
	/* 通用状态码 */

	// SUCCESS 成功标识
	SUCCESS = 0
	// 请求成功
	StatusOK = 200
	// 请求已被实现
	RequestCompleted = 201
	// 请求参数有误
	StatusBadRequest = 400
	// 请求参数缺少有效的身份验证凭据
	StatusUnauthorized = 401
	// ErrorTokenWrong token错误
	ErrorTokenWrong = 1001
	// ErrorTokenFormatWrong token格式错误
	ErrorTokenFormatWrong = 1002
	// ErrorUserNameFormatWrong 用户名格式错误
	ErrorUserNameFormatWrong = 1003
	// ERROR 错误标识
	ERROR = 9999

	/* 业务状态码 */

	//用户创建失败
	ErrorUserCreatedFail = 10001
	//找不到对应用户信息
	ErrorUserNotFound = 10002
	//用户信息更新失败
	ErrorUserUpdatedFail = 10003
	//项目创建失败
	ErrorProjectCreatedFail = 10101
	//找不到对应项目信息
	ErrorProjectNotFound = 10102
	//版本创建失败
	ErrorReleaseCreatedFail = 10201
	//找不到对应版本信息
	ErrorReleaseNotFound = 10202
)

// 错误码映射表
var codeMap = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ErrorTokenWrong:          "Token 不正确",
	ErrorTokenFormatWrong:    "Token格式错误",
	ErrorUserNameFormatWrong: "用户名格式错误",
	StatusOK:                 "请求成功",
	RequestCompleted:         "请求已被实现",
	StatusBadRequest:         "请求参数有误",
	StatusUnauthorized:       "请求参数缺少有效的身份验证凭据",

	ErrorUserCreatedFail:    "用户创建失败",
	ErrorUserNotFound:       "找不到对应用户信息",
	ErrorUserUpdatedFail:    "用户信息更新失败",
	ErrorProjectCreatedFail: "项目创建失败",
	ErrorReleaseNotFound:    "找不到对应版本信息",
}

// GetErrMsg 获取错误信息
func GetErrMsg(code int) string {
	return codeMap[code]
}

/* type Response struct {
	code    int
	message error
	data    interface{}
}

var (
	//用户创建失败
	ErrUserCreatedFailed = errors.New("User creation failed")
	//找不到对应用户信息
	ErrUserNotFound = errors.New("The corresponding user information cannot be found")
	//用户信息更新失败
	ErrUserUpdatedFailed = errors.New("User information update failed")
	//项目创建失败
	ErrProjectCreatedFailed = errors.New("Project creation failed")
	//找不到对应项目信息
	ErrProjectNotFound = errors.New("The corresponding item information cannot be found")
	//版本创建失败
	ErrReleaseCreatedFailed = errors.New("Release creation failed")
	//找不到对应版本信息
	ErrReleaseNotFound = errors.New("The corresponding release information cannot be found")
) */
