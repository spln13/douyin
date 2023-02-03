package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserLogin struct {
	ID                  int64  `gorm:"primary_key"`
	Username            string `gorm:"primary_key"`
	Password            string
	Token               string
	CreateTime          time.Time
	UpdateTime          time.Time
	TokenExpirationTime time.Time
}

// CheckUsernameUnique
// 检查该用户名是否唯一
func (u *UserLogin) CheckUsernameUnique() bool {
	var user UserInfo
	err := GetDB().Where("username = ?", u.Username).Find(&user).Error
	if err != nil {
		log.Println(err)
	}
	if user.ID == 0 {
		return true
	}
	return false
}

// CheckParamValid
// 检查用户名以及密码是否符合规范
func (u *UserLogin) CheckParamValid(MaxUsernameLength int) error {
	if u.Username == "" || len(u.Username) > MaxUsernameLength {
		return errors.New("用户名长度不符合规范")
	}
	return nil
}

// SaveUser
// 向数据库中插入新用户，若出错则回滚
func (u *UserLogin) SaveUser() error {
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		// 进行插入事务
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		userInfo := &UserInfo{
			UserID:        u.ID,
			FollowCount:   0,
			FollowerCount: 0,
			UpdateTime:    time.Now(),
		}
		if err := tx.Create(userInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateUser
// 根据主键更新数据库中用户数据
func (u *UserLogin) UpdateUser() error {
	// TODO: 待完成
	return nil
}

var signature = "douyinSignature"

type Claims struct {
	ID                 int64
	jwt.StandardClaims // jwt中标准格式,主要是设置token的过期时间
}

// GenerateToken
// 调用库的NewWithClaims(加密方式,载荷).SignedString(签名) 生成token
func (u *UserLogin) GenerateToken() {
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
	u.TokenExpirationTime = expirationTime
}

// QueryByUsername
// 查询用户名对应的用户
func QueryByUsername(username string) *UserLogin {
	var user UserLogin
	err := GetDB().Where("username = ?", username).Omit("create_time", "update_time").First(&user).Error
	if err != nil {
		log.Println(err)
	}
	log.Println(user) // TEST
	return &user
}
