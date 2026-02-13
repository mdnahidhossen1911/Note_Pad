package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxRequests    = 6
	windowDuration = time.Second
	blockDuration  = 5 * time.Minute
)

type clientInfo struct {
	requests int
	window   time.Time
	blocked  time.Time
}

var (
	clients = make(map[string]*clientInfo)
	mu      sync.Mutex
)

// RateLimiter limits each IP to 6 req/sec; blocks for 5 min on abuse.
func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		defer mu.Unlock()

		info, ok := clients[ip]
		if !ok {
			info = &clientInfo{window: time.Now()}
			clients[ip] = info
		}

		// Still blocked?
		if !info.blocked.IsZero() && time.Now().Before(info.blocked) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests — your IP is blocked for 5 minutes.",
			})
			c.Abort()
			return
		}

		// New window?
		if time.Since(info.window) > windowDuration {
			info.requests = 0
			info.window = time.Now()
		}

		info.requests++
		if info.requests > maxRequests {
			info.blocked = time.Now().Add(blockDuration)
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests — your IP is blocked for 5 minutes.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
