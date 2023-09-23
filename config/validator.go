package config

import (
	utils "Cgo/utils/validator"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	// 初始化验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 将 json tag 作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			// 如果 json tag 为 - 则不处理
			if name == "-" {
				return ""
			}
			return name
		})
	}

}
