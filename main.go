package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"learn-go/router"
	"learn-go/util"
)

//test
func main() {

	//在这里自定义jwt的密钥
	var secretKey = []byte("zmb")

	type myClaim struct {
		jwt.StandardClaims
	}

	var claim = jwt.StandardClaims{
		ExpiresAt: 60 * 60, //1小时后过期
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	test, _ := token.SignedString(secretKey)
	fmt.Println(test)
	fmt.Println(token)

	//加载配置
	util.LoadConfig()
	//连接数据库
	util.ConnectDatabase()
	//开始采用自定义的方式生成引擎
	engine := router.InitRouter()
	err := engine.Run(":" + util.MyConfig.HttpPort)
	if err != nil {
		panic("端口错误或未知错误...")
	}

}
