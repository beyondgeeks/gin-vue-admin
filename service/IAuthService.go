package service

//IUserService UserService接口定义
type IUserService interface {
	//CheckUser 身份验证
	CheckUser(username string, password string) bool
	//GetUserAvatar 获取用户头像
	GetUserAvatar(username string) string
	//GetRoles 获取用户角色
	GetRoles(username string) []string
}
