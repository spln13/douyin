package service

import (
	"douyin/middlewares"
	"douyin/models"
	"errors"
)

const (
	MaxUsernameLength = 30
)

type UserRegisterLoginFlow struct {
	ID       int64
	Username string
	Password string
	Token    string
}

func NewUserRegisterLoginFlow(username string, password string) *UserRegisterLoginFlow {
	return &UserRegisterLoginFlow{
		Username: username,
		Password: password,
	}
}

func (u *UserRegisterLoginFlow) DoRegister() error {
	if err := u.CheckParamValid(MaxUsernameLength); err != nil { // 检查用户名格式是否符合规范
		return err
	}
	userLoginDAO := &models.UserLogin{Username: u.Username, Password: u.Password}
	if err := userLoginDAO.CheckUsernameUnique(); err != nil { // 检查用户名是否存在
		return err
	}
	if err := userLoginDAO.SaveUser(); err != nil {
		return err
	} //向数据库中插入用户
	u.ID = userLoginDAO.ID // 向前端返回参数需要插入用户的id
	var err error
	u.Token, err = middlewares.ReleaseToken(userLoginDAO) // 根据用户id生成token
	return err
}

// CheckParamValid
// 检查用户名以及密码是否符合规范
func (u *UserRegisterLoginFlow) CheckParamValid(MaxUsernameLength int) error {
	if u.Username == "" || len(u.Username) > MaxUsernameLength {
		return errors.New("用户名长度不符合规范")
	}
	return nil
}

func (u *UserRegisterLoginFlow) DoLogin() error {
	userLoginDao := models.QueryByUsername(u.Username)
	if u.ID = userLoginDao.ID; u.ID == 0 {
		return errors.New("用户不存在")
	}
	if u.Password != userLoginDao.Password {
		return errors.New("密码错误")
	}
	var err error
	if u.Token, err = middlewares.ReleaseToken(userLoginDao); err != nil {
		return err
	}
	return nil
}
