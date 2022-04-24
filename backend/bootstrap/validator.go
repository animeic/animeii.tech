package bootstrap

import (
	"animeic-gin/app/request"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	request
	// 自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", request.ValidateMobile)
		_ = v.RegisterValidation("email", request.ValidateEmail)
		_ = v.RegisterValidation("username", request.ValidateUserName)
		_ = v.RegisterValidation("password", request.ValidatePassword)

		// 注册自定义json tag函数
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 这里有空格号分割
			name := strings.SplitN(field.Tag.Get("json"), " ", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
