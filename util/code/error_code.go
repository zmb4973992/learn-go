package code

const (
	Success int = 0
	Error   int = 500

	Error_Record_Not_Found int = 1001

	//用户模块的错误
	Error_Username_Or_Password_Exist int = 2001
	Error_Username_Exist             int = 2002
	Error_Password_Incorrect         int = 2003

	//JWT鉴权的错误
	Error_Token_Invalid int = 3001
)

var ErrorMessage = map[int]string{
	Success: "成功",
	Error:   "错误",

	Error_Record_Not_Found: "未找到指定记录",

	Error_Username_Or_Password_Exist: "用户名或密码错误",
	Error_Username_Exist:             "用户名已存在",
	Error_Password_Incorrect:         "密码错误",
	Error_Token_Invalid:              "token无效",
}

func GetErrorMessage(code int) string {
	message, ok := ErrorMessage[code]
	if !ok {
		return ErrorMessage[Error]
	}
	return message
}
