package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fastrand"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
}

func NewApp(db *gorm.DB) *App {
	return &App{db: db}
}

func (a *App) Run() error {
	app := fiber.New()

	// Setup routes
	app.Get("/", a.handleIndex)

	// Start server
	log.Info("Server is running on http://127.0.0.1:3000")
	return app.Listen(":3000")
}

func (a *App) RunGin() error {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", a.handleGinIndex)

	log.Info("Server is running on http://127.0.0.1:3000")
	return r.Run(":3000")
}

func (a *App) handleIndex(c fiber.Ctx) error {
	randomID1 := int32(fastrand.Uint32n(1000))
	randomID2 := int32(fastrand.Uint32n(1000)) + 8000
	slog.Infof("Generated random id1: %d, id2: %d", randomID1, randomID2)

	var users []SysUser
	tx := a.db.Where("user_id BETWEEN ? AND ?", randomID1, randomID2).Find(&users)
	if tx.Error != nil {
		log.Error("Failed to query users", zap.Error(tx.Error))
		return c.Status(fiber.StatusOK).JSON([]SysUser{})
	}
	slog.Infof("Found %d users", len(users))

	return c.Status(fiber.StatusOK).JSON(users)
}

func (a *App) handleGinIndex(c *gin.Context) {
	randomID1 := int32(fastrand.Uint32n(1000))
	randomID2 := int32(fastrand.Uint32n(1000)) + 8000
	slog.Infof("Generated random id1: %d, id2: %d", randomID1, randomID2)

	var users []SysUser
	tx := a.db.Where("user_id BETWEEN ? AND ?", randomID1, randomID2).Find(&users)
	if tx.Error != nil {
		log.Error("Failed to query users", zap.Error(tx.Error))
		c.JSON(fiber.StatusOK, []SysUser{})
	}
	slog.Infof("Found %d users", len(users))

	c.JSON(http.StatusOK, users)
}
