package jwt

import (
	"github.com/golang-jwt/jwt"
	"learn-go/util"
	"time"
)

type myClaim struct {
	Username string
	jwt.StandardClaims
}

// GenerateToken 传入Username，返回token字符串
func GenerateToken(Username string) string {
	claim := myClaim{
		Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), //默认7天的有效期
		}}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//fmt.Println(util.MyJWTConfig.SecretKey)  //用于测试密钥是否正常
	tokenString, _ := tokenStruct.SignedString(util.MyJWTConfig.SecretKey)
	return tokenString
}

// ParseToken 验证用户token。这部分基本就是参照官方写法。
//第一个参数是token字符串，第二个参数是结构体，第三个参数是jwt规定的解析函数，包含密钥
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
