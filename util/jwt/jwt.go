package jwt

import (
	"github.com/golang-jwt/jwt"
	"learn-go/util"
	"time"
)

//在这里自定义jwt的密钥
//var secretKey = util.MyJWTConfig.SecretKey

type myClaim struct {
	UserID int64
	jwt.StandardClaims
}

// GenerateToken 传入UserID，返回token字符串或者错误信息
func GenerateToken(UserID int64) (string, error) {
	claim := myClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(util.MyJWTConfig.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 验证用户token。第一个参数是token字符串，第二个参数是结构体，第三个参数是jwt规定的解析函数，包含密钥
func ParseToken(token string) (*myClaim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &myClaim{}, func(token *jwt.Token) (interface{}, error) {
		return util.MyJWTConfig.SecretKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*myClaim); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
