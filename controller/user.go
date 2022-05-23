package controller

import (
	"github.com/gin-gonic/gin"
	"learn-go/model"
	"learn-go/util/jwt"
	"net/http"
	"strconv"
)

// CreateUser test
func CreateUser(context *gin.Context) {
	var user model.User
	_ = context.ShouldBind(&user)

}

func GenerateToken(ctx *gin.Context) {
	userID, _ := strconv.ParseInt(ctx.Param("userid"), 10, 64)
	token, _ := jwt.GenerateToken(userID)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func ParseToken(c *gin.Context) {

	token1 := c.GetHeader("authorization")
	c.Header("authorization", "ok111")
	token2 := c.GetHeader("authorization")
	c.JSON(200, gin.H{
		"t1": token1,
		"t2": token2,
	})
	//myClaim, code := jwt.ParseToken(token)
	//if code != nil {
	//	return
	//}
	//userID := myClaim.UserID
	//expireAt := myClaim.ExpiresAt
	//c.JSON(http.StatusOK, gin.H{
	//	"用户id": userID,
	//	"过期时间": time.Unix(expireAt, 0).Format("2006-01-02 15:04:05"),
	//	"当前时间": time.Now().Format("2006-01-02 15:04:05"),
	//})
}
