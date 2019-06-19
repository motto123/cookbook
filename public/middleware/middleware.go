package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var skipUrlForTokenArr = []string{
	"/version",
	"/public/",
	"/user/test",
	"/user/token/refresh",
	"/user/wxserver/connect",
	"/recharge/callback",
}

var maybeCheckForTokenArr = []string{
	"/user/login",
	"/user/register",
}

var skipParamsForSQLInjectMap = map[string]int{
	"files":             1,
	"file":              1,
	"Identifier":        1,
	"name":              1,
	"link":              1,
	"head_pic_link":     1,
	"pics":              1,
	"wechat_qr_code":    1,
	"id_front_pic":      1,
	"id_reverse_pic":    1,
	"car_front_pic":     1,
	"car_reverse_pic":   1,
	"house_front_pic":   1,
	"house_reverse_pic": 1,
	"education_pic":     1,
	"head_pic":          1,
	"back_pic":          1,
	"encrypt_data":      1,
}

type User struct {
	UUid     string
	Nickname string
	Sex      int
	Status   int
	Role     int
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()

	}
}
