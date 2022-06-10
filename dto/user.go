package dto

import "learn-go/model"

type UserDTO struct {
	model.BaseModel
	Username string
}
