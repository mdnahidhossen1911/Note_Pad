package routes

import (
	"note_pad/config"
	notecontroller "note_pad/controllers/note_controller"
	usercontroller "note_pad/controllers/user_controller"
	"note_pad/middleware"
	"note_pad/repositories"
	noteservice "note_pad/services/note_service"
	userService "note_pad/services/user_service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {

	userRepo := repositories.NewUserRepository(db)
	noteRepo := repositories.NewNoteRepository(db)

	userSvc := userService.NewUserService(userRepo, cfg)
	noteSvc := noteservice.NewNoteService(cfg.JwtSecureKey, noteRepo)

	userCtrl := usercontroller.NewUserController(userSvc)
	noteCtrl := notecontroller.NewNoteController(noteSvc)

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
	registerNoteRoutes(api, noteCtrl, userRepo, cfg.JwtSecureKey)
	return r
}
