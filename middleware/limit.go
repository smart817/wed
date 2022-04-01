package middleware

//接口限流
import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

//Limiter 限流器对象
type limiter struct {
	value int64
	max   int64
	ts    int64
}

//NewLimiter 产生一个限流器
func newLimiter(cnt int64) *limiter {
	return &limiter{
		value: 0,
		max:   cnt,
		ts:    time.Now().Unix(),
	}
}

//Ok 是否可以通过
func (l *limiter) Ok() bool {
	ts := time.Now().Unix()
	tsOld := atomic.LoadInt64(&l.ts)
	if ts != tsOld {
		atomic.StoreInt64(&l.ts, ts)
		atomic.StoreInt64(&l.value, 1)
		return true
	}
	return atomic.AddInt64(&(l.value), 1) < l.max
}

//SetMax 设置最大限制
// func (l *limiter) SetMax(m int64) {
// 	l.max = m
// }

//MaxAllowed 限流器 r.Use(MaxAllowed(200)) 限制每秒最多允许200个请求
func MaxAllowed(limitValue int64) func(c *gin.Context) {
	limiter := newLimiter(limitValue)
	//log.Info("limiter.SetMax:", limitValue)
	// 返回限流逻辑
	return func(c *gin.Context) {
		if !limiter.Ok() {
			c.AbortWithStatus(http.StatusServiceUnavailable) //超过每秒200，就返回503错误码
			return
		}
		c.Next()
	}
}
