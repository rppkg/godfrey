package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// TokenBucket 结构体用于限流.
type TokenBucket struct {
	capacity  int        // 令牌桶的容量
	tokens    int        // 当前令牌数
	rate      int        // 令牌生成速率（每秒生成多少个令牌）
	mutex     sync.Mutex // 互斥锁，保护令牌桶的并发访问
	lastToken time.Time  // 上次生成令牌的时间
}

// NewTokenBucket 创建一个新的令牌桶; 指定容量数和每秒生成令牌数.
func NewTokenBucket(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		capacity:  capacity,
		tokens:    capacity,
		rate:      rate,
		lastToken: time.Now(),
	}
}

// Allow 方法检查是否允许一个请求通过.
func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastToken).Seconds()
	tb.tokens += int(elapsed * float64(tb.rate))
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastToken = now

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

// RateLimitMiddleware 限流中间件.
func RateLimitMiddleware(tb *TokenBucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !tb.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}
		c.Next()
	}
}
