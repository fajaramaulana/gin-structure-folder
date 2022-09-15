package utils

import "github.com/go-playground/validator/v10"

type response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatError(err error) []string {
	var errors []string
	for _, v := range err.(validator.ValidationErrors) {
		errors = append(errors, v.Error())
	}

	return errors
}
