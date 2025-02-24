package main

import (
	sys_log "log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	newLogger := logger.New(
		sys_log.New(os.Stdout, "\r\n", sys_log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             10 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Silent,    // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,             // Don't include params in the SQL log
			Colorful:                  false,            // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open("host=127.0.0.1 user=postgres password=12345678 dbname=postgres port=5432 sslmode=disable search_path=public TimeZone=Asia/Shanghai"), &gorm.Config{
		Logger:                 newLogger,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	sqlDB.SetConnMaxLifetime(time.Hour * 8)

	return db
}
