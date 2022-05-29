package status

const (
	Success int = 0
	Error   int = 500

	// ErrorRecordNotFound 未找到指定记录
	ErrorRecordNotFound int = 1001
	// ErrorNotEnoughParameters 没有足够的参数
	ErrorNotEnoughParameters int = 1002
	// ErrorInvalidURIParameters URI参数无效
	ErrorInvalidURIParameters int = 1003
	// ErrorInvalidFormDataParameters form-data参数无效
	ErrorInvalidFormDataParameters int = 1004
	// ErrorFailToSaveRecord 保存记录失败
	ErrorFailToSaveRecord int = 1005
	// ErrorFailToDeleteRecord 删除记录失败
	ErrorFailToDeleteRecord int = 1006

	ErrorUsernameOrPasswordExist int = 2001
	ErrorUsernameExist           int = 2002
	ErrorPasswordIncorrect       int = 2003

	ErrorTokenInvalid int = 3001
)

var Message = map[int]string{
	Success: "成功",
	Error:   "错误",

	ErrorRecordNotFound:            "未找到指定记录",
	ErrorNotEnoughParameters:       "没有足够的参数",
	ErrorInvalidURIParameters:      "URI参数无效",
	ErrorInvalidFormDataParameters: "form-data参数无效",
	ErrorFailToSaveRecord:          "保存记录失败",
	ErrorFailToDeleteRecord:        "删除记录失败",

	ErrorUsernameOrPasswordExist: "用户名或密码错误",
	ErrorUsernameExist:           "用户名已存在",
	ErrorPasswordIncorrect:       "密码错误",
	ErrorTokenInvalid:            "token无效",
}

func GetMessage(code int) string {
	message, ok := Message[code]
	if !ok {
		return "获取错误信息失败，请检查status/code_and_message相关配置"
	}
	return message
}
