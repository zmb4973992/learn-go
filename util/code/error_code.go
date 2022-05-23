package code

const (
	Success int = 0
	Error   int = 500

	//用户模块的错误
	Error_Username_Exist     int = 1001
	Error_Password_Incorrect int = 1002
)

var ErrorMessage = map[int]string{
	Success:                  "",
	Error:                    "发生错误",
	Error_Username_Exist:     "用户名已存在",
	Error_Password_Incorrect: "密码错误",
}

func GetErrorMessage(code int) string {
	message, ok := ErrorMessage[code]
	if !ok {
		return ErrorMessage[Error]
	}
	return message
}
