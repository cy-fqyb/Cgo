package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^1([358][0-9]|14[579]|5[^4]16[6]|7[1-35-8]|9[189])\d{8}$`, mobile) // 用正则去匹配
	return ok
}
