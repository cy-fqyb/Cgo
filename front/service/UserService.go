package service

type userService struct{}

var UserService = userService{}

func (userService) Login() string {
	return "前端页面登录成功"
}
func (userService) Register() string {
	return "用户注册接口"
}
