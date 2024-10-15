package service

import "github.com/go-playground/validator/v10"

var validate = validator.New(validator.WithRequiredStructEnabled())

type Page struct {
	PageNum  int `json:"page_num" validate:"required,page_num" default:"1"`
	PageSize int `json:"page_size" validate:"required,page_num" default:"10"`
}
