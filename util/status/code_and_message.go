package status

//自定义错误的code
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
	// ErrorInvalidJsonParameters json参数无效
	ErrorInvalidJsonParameters int = 1005
	// ErrorInvalidQueryParameters query参数无效
	ErrorInvalidQueryParameters int = 1006
	// ErrorFailToSaveRecord 保存记录失败，请重试
	ErrorFailToSaveRecord int = 1010
	// ErrorFailToDeleteRecord 删除记录失败，请重试
	ErrorFailToDeleteRecord int = 1011
	ErrorFileTooLarge       int = 1012
	ErrorFailToUploadFiles  int = 1013

	ErrorInvalidUsernameOrPassword int = 2001
	ErrorUsernameExist             int = 2002
	ErrorPasswordIncorrect         int = 2003

	ErrorAccessTokenInvalid  int = 3001
	ErrorAccessTokenNotFound int = 3002
	ErrorPermissionDenied    int = 3100
	ErrorNeedAdminPrivilege  int = 3101

	ErrorFailToEncrypt  int = 4001
	ErrorInvalidRequest int = 4002
)

// Message 自定义错误的message
var Message = map[int]string{
	Success: "成功",
	Error:   "错误",

	ErrorRecordNotFound:            "未找到指定记录",
	ErrorNotEnoughParameters:       "没有足够的参数",
	ErrorInvalidURIParameters:      "URI参数无效",
	ErrorInvalidFormDataParameters: "form-data参数无效",
	ErrorInvalidJsonParameters:     "json参数缺失或json参数无效",
	ErrorInvalidQueryParameters:    "query参数无效",
	ErrorFailToSaveRecord:          "保存记录失败",
	ErrorFailToDeleteRecord:        "删除记录失败",
	ErrorFileTooLarge:              "文件过大",
	ErrorFailToUploadFiles:         "上传文件失败",

	ErrorInvalidUsernameOrPassword: "用户名或密码错误",
	ErrorUsernameExist:             "用户名已存在",
	ErrorPasswordIncorrect:         "密码错误",

	ErrorAccessTokenInvalid:  "access_token无效",
	ErrorAccessTokenNotFound: "缺少access_token",
	ErrorPermissionDenied:    "权限不足",
	ErrorNeedAdminPrivilege:  "权限不足，该操作需要管理员权限",

	ErrorFailToEncrypt:  "加密失败",
	ErrorInvalidRequest: "请求路径错误",
}

func GetMessage(code int) string {
	message, ok := Message[code]
	if !ok {
		return "由于错误代码未定义返回信息，导致获取错误信息失败，建议检查status/code_and_message相关配置"
	}
	return message
}
