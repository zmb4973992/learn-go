package serializer

type CommonResponse struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserLoginResponse struct {
	Username string
	Token    string
}
