package request

// type Register struct {
// 	Etype    string `form:"etype" json:"etype" binding:"required"`
// 	Ecode    string `form:"ecode" json:"ecode" binding:"required,min=6,max=6"`
// 	Username string `form:"username" json:"username" binding:"required,username"`
// 	Email    string `form:"email" json:"email" binding:"required,email"`
// 	Password string `form:"password" json:"password" binding:"required,password"`
// }

// func (register Register) GetMessages() ValidatorMessages {
// 	return ValidatorMessages{
// 		"etype.required":    "验证类型不能为空",
// 		"ecode.required":    "昵称不能为空",
// 		"username.required": "账号不能为空",
// 		"username.username": "账号4-16字母数字下划线中划线组成",
// 		"email.required":    "邮箱不能为空",
// 		"email.email":       "邮箱格式不正确",
// 		"password.required": "用户密码不能为空",
// 		"password.password": "密码是6-20位字母数字",
// 	}
// }
