package jwts

import (
	"errors"
	"fast_gin/global"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	UserID uint `json:"userID"`
	RoleID int8 `json:"roleID"`
}

type MyClaims struct {
	Claims
	jwt.RegisteredClaims
}

// 生成token
func SetToken(data Claims) (string, error) {
	SetClaims := MyClaims{
		Claims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), //有效时间
			Issuer:    global.Config.Jwt.Issuer,                                                                 //签发人

		},
	}
	//使用指定的加密方式和声明类型创建新令牌，此时将数据跟算法组装好（head.payload)
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodES256, SetClaims)
	//获取完整的，签名令牌(head.payload.signature)
	token, err := tokenStruct.SignedString([]byte(global.Config.Jwt.Key))
	if err != nil {
		logrus.Errorf("颁发jwt失败 %s", err)
		return "", err
	}
	return token, nil
}

// CheckToken 验证Token
func CheckToken(token string) (*MyClaims, error) {
	//解析，验证并返回token
	tokenObj, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(global.Config.Jwt.Key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenObj.Claims.(*MyClaims); ok && tokenObj.Valid {
		return claims, nil

	} else {
		return nil, errors.New("token无效")
	}
}
