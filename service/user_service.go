package service

import (
	"errors"
	"ginchat/common"
	"ginchat/common/request"
	"ginchat/model"
	"ginchat/util"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user model.User) {
	var result = common.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&model.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = model.User{Name: params.Name, Mobile: params.Mobile, Password: util.BcryptMake([]byte(params.Password))}
	err = common.App.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *model.User) {
	err = common.App.DB.Where("mobile = ?", params.Mobile).First(&user).Error
	if err != nil || !util.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user model.User) {
	intId, err := strconv.Atoi(id)
	err = common.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
func (jwtService *jwtService) GetUserInfo(GuardName string, id string) (err error, user JwtUser) {
	switch GuardName {
	case AppGuardName:
		return UserService.GetUserInfo(id)
	default:
		err = errors.New("guard " + GuardName + " does not exist")
	}
	return
}
