package resp

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{"code": 200, "msg": msg, "data": data})
}

func Info(c *gin.Context, msg string) {
	c.JSON(200, gin.H{"code": 200, "msg": msg})
}

func Error(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{"code": 1, "msg": msg})
}

//200 正常，1，参数错误，2 未知错误
