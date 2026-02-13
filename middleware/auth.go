package middleware

import (
	"fmt"
	"net/http"
	"note_pad/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const AuthUserKey = "authUser"

// AuthRequired validates the Bearer JWT token and sets user payload in context.
func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format: Bearer <token>"})
			c.Abort()
			return
		}

		payload, err := utils.VerifyJWT(parts[1], jwtSecret)
		if err != nil {
			errMsg := "Invalid token"
			if err.Error() == "token expired" {
				errMsg = "Token expired. Please login again."
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			c.Abort()
			return
		}

		// Calculate remaining access time
		expiresAt := time.Unix(payload.Exp, 0)
		remaining := time.Until(expiresAt)

		// Add token info to response headers
		c.Header("X-Token-Expires-At", expiresAt.UTC().Format(time.RFC3339))
		c.Header("X-Token-Remaining", formatDuration(remaining))

		c.Set(AuthUserKey, payload)
		c.Next()
	}
}

// formatDuration formats duration in human-readable format
func formatDuration(d time.Duration) string {
	if d < 0 {
		return "expired"
	}
	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}

// GetAuthUser retrieves the authenticated user payload from context.
func GetAuthUser(c *gin.Context) *utils.JWTPayload {
	v, _ := c.Get(AuthUserKey)
	if p, ok := v.(*utils.JWTPayload); ok {
		return p
	}
	return nil
}
