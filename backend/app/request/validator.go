package request

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)
		for _, v := range err.(validator.ValidationErrors) {
			if isValidator {
				if message, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return message
				}
			}
			return v.Error()
		}
	}
	return "Parameter error"
}

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	return ok
}

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	ok, _ := regexp.MatchString(`^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$`, email)
	return ok
}
func ValidateUserName(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_-]{4,16}$`, username)
	return ok
}
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	// 	密码至少包含 数字和英文，长度6-20
	//    String reg = "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,20}$";
	// 	密码包含 数字,英文,字符中的两种以上，长度6-20
	//    String reg = "^(?![0-9]+$)(?![a-z]+$)(?![A-Z]+$)(?!([^(0-9a-zA-Z)])+$).{6,20}$";
	// 	至少包含数字跟字母，可以有字符
	//    String reg = "(?=.*([a-zA-Z].*))(?=.*[0-9].*)[a-zA-Z0-9-*/+.~!@#$%^&*()]{6,20}$";
	ok, _ := regexp.MatchString(`[0-9A-Za-z]{6,20}$`, password)
	return ok
}
