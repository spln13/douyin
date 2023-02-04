package service

import (
	"douyin/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
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
	userLoginDAO := models.UserLogin{Username: u.Username, Password: u.Password}
	if err := userLoginDAO.CheckUsernameUnique(); err != nil { // 检查用户名是否存在
		return err
	}
	userLoginDAO.Token = u.GenerateToken()                                // 生成token
	userLoginDAO.TokenExpirationTime = time.Now().Add(7 * 24 * time.Hour) // 设置token有效时间为一周
	userLoginDAO.CreateTime = time.Now()
	userLoginDAO.UpdateTime = time.Now()
	err := userLoginDAO.SaveUser() // 将用户插入到数据库
	u.ID = userLoginDAO.ID         // 向前端返回参数需要插入用户的id
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
		log.Println(u.Password, userLoginDao.Password)
		return errors.New("密码错误")
	}
	userLoginDao.Token = u.GenerateToken() // 重新生成token
	userLoginDao.TokenExpirationTime = time.Now().Add(7 * 24 * time.Hour)
	if err := userLoginDao.UpdateUserToken(); err != nil {
		return errors.New("系统错误，请重新登陆")
	}
	return nil
}

// 生成token
var signature = "douyinSignature"

type Claims struct {
	ID                 int64
	jwt.StandardClaims // jwt中标准格式,主要是设置token的过期时间
}

// GenerateToken
// 调用库的NewWithClaims(加密方式,载荷).SignedString(签名) 生成token
func (u *UserRegisterLoginFlow) GenerateToken() string {
	nowTime := time.Now()
	expirationTime := nowTime.Add(7 * 24 * time.Hour) // 过期时间
	issuer := "linan"
	claims := Claims{
		ID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 转成纳秒
			Issuer:    issuer,
		},
	}
	// 根据签名生成token，NewWithClaims(加密方式,claims) ==》 头部，载荷，签证
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signature))
	if err != nil {
		log.Println(err)
	}
	u.Token = token
	return token
}
