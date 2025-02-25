package main

import (
	"go.uber.org/zap"
)

func main() {
	if err := startServer(); err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}
}

func startServer() error {
	db := ConnectDB()
	log.Info("Connected to database successfully")

	app := NewApp(db)
	return app.RunGin()
}
