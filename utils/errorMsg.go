package utils

// 业务错误码

const (
	// SUCCESS 成功标识
	SUCCESS = 0
	// ERROR 错误标识
	ERROR = 9999
	// ErrorTokenWrong token错误
	ErrorTokenWrong = 1001
	// ErrorTokenFormatWrong token格式错误
	ErrorTokenFormatWrong = 1002
	// ErrorUserNameFormatWrong 用户名格式错误
	ErrorUserNameFormatWrong = 1003
)

// 错误码映射表
var codeMap = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ErrorTokenWrong:          "Token 不正确",
	ErrorTokenFormatWrong:    "Token格式错误",
	ErrorUserNameFormatWrong: "用户名格式错误",
}

// GetErrMsg 获取错误信息
func GetErrMsg(code int) string {
	return codeMap[code]
}
