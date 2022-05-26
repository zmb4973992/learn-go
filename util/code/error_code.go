package code

const (
	Success int = 0
	Error   int = 500

	// ErrorRecordNotFound 未找到指定记录
	ErrorRecordNotFound int = 1001
	// ErrorNotEnoughParameters 没有足够的参数
	ErrorNotEnoughParameters int = 1002

	ErrorUsernameOrPasswordExist int = 2001
	ErrorUsernameExist           int = 2002
	ErrorPasswordIncorrect       int = 2003

	ErrorTokenInvalid int = 3001
)

var ErrorMessage = map[int]string{
	Success: "成功",
	Error:   "错误",

	ErrorRecordNotFound:      "未找到指定记录",
	ErrorNotEnoughParameters: "没有足够的参数",

	ErrorUsernameOrPasswordExist: "用户名或密码错误",
	ErrorUsernameExist:           "用户名已存在",
	ErrorPasswordIncorrect:       "密码错误",
	ErrorTokenInvalid:            "token无效",
}

func GetErrorMessage(code int) string {
	message, ok := ErrorMessage[code]
	if !ok {
		return ErrorMessage[Error]
	}
	return message
}
