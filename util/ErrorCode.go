package util

const (
	Success = 200
	Error   = 500

	//用户模块的错误
	Error_Username_Exist     = 1001
	Error_Password_Incorrect = 1002
)

var ErrorMessage = map[uint64]string{
	Success:                  "ok",
	Error:                    "发生错误",
	Error_Username_Exist:     "用户名已存在",
	Error_Password_Incorrect: "密码错误",
}

func GetErrorMessage(code uint64) string {
	return ErrorMessage[code]
}
