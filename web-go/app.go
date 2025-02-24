package main

import (
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
	app.Get("/", func(c fiber.Ctx) error {
		return a.handleIndex(c)
	})

	// Start server
	log.Info("Server is running on http://127.0.0.1:3000")
	return app.Listen(":3000")
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
