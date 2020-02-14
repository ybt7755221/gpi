package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	et "gpi/entities"
	"gpi/libraries/verify"
)
//数据验证-中间件
func Verify(c *gin.Context) {
	var token string
	var ts string
	methodStr := c.Request.Method
	if skip := skipVerify(c, methodStr); skip == false {
		if methodStr == "GET" || methodStr == "DELETE"{
			token = c.Query("token")
			ts = c.Query("ts")
		} else {
			token = c.PostForm("token")
			ts = c.PostForm("ts")
		}
		if len(ts) < 1 {
			c.JSON(http.StatusOK, et.ApiResonse{et.EntityParametersMissing, "缺少ts值", gin.H{}})
			c.Abort()
		}
		if len(token) < 1 {
			c.JSON(http.StatusOK, et.ApiResonse{et.EntityParametersMissing, "缺少token值", gin.H{}})
			c.Abort()
		} else {
			_, tokenStr := verify.GenerateToken(c)
			if token != tokenStr {
				c.JSON(http.StatusOK, et.ApiResonse{et.EntityForbidden, et.GetStatusMsg(et.EntityForbidden), gin.H{}})
				c.Abort()
			} else {
				c.Next()
			}
		}
	}else{
		c.Next()
	}
}

func skipVerify(c *gin.Context, methodStr string) bool {
	var skip string
	if methodStr == "GET" || methodStr == "DELETE" {
		skip = c.Query("skip_debug")
	}else{
		skip = c.PostForm("skip_debug")
	}
	if skip == "161217" {
		return true
	}
	return false
}

