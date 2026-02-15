package routes

import (
	"note_pad/config"
	"note_pad/controllers"
	"note_pad/middleware"
	"note_pad/repositories"
	userService "note_pad/services/user_service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter wires up all dependencies and registers all routes.
// This is the composition root — Model ↔ Service ↔ Controller ↔ Route.
func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {

	// ── Repositories (Model / Data layer) ───────────────────────────────
	userRepo := repositories.NewUserRepository(db)
	noteRepo := repositories.NewNoteRepository(db)

	if noteRepo != nil {
		// wrte condition
	}

	// ── Services (Business logic layer) ─────────────────────────────────
	userSvc := userService.NewUserService(userRepo, cfg)

	// ── Controllers (C in MVC) ───────────────────────────────────────────
	userCtrl := controllers.NewUserController(userSvc)

	// ── Gin Engine ───────────────────────────────────────────────────────
	r := gin.New()

	// Global middleware
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimiter())
	r.Use(gin.Recovery())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": cfg.ServiceName,
			"version": cfg.Version,
		})
	})

	// ── API v1 ───────────────────────────────────────────────────────────
	api := r.Group("/api/v1")

	registerUserRoutes(api, userCtrl, userRepo, cfg.JwtSecureKey)
	return r
}
