package dto

type UpdateProfileDto struct {
	Name    string `json:"name" validate:"required,min=1,max=100"`
	Picture string `json:"picture" validate:"required,min=1,max=100"`
}
