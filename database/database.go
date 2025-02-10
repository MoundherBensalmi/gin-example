package database

import (
	"MBFacto/config"
	"MBFacto/utils/log_colors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func ConnectToDB() {
	maxRetries := 3
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.Database.Username,
		config.Cfg.Database.Password,
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.Name,
	)

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		// if connected successfully, break the loop
		if err == nil {
			break
		}

		log_colors.CLog('r', "Error connecting to the database.")
		log_colors.CLog('b', "Attempt", i+1, "- Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	// if not connected after maxRetries, log a fatal error
	if err != nil {
		log_colors.CFLog('r', "Failed to connect to the database after ", maxRetries, " attempts.")
	}

	// Access the underlying *sql.DB object to configure the connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log_colors.CFLog('r', "error getting *sql.DB object:", err)
	}

	// Set the connection pool configuration
	sqlDB.SetMaxOpenConns(100)                 // Maximum number of open connections
	sqlDB.SetMaxIdleConns(10)                  // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // Maximum connection lifetime

	log_colors.CLog('g', "Connected to the database with connection pooling.")
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log_colors.CLog('r', "Error getting *sql.DB object:", err)
			return
		}

		err = sqlDB.Close()
		if err != nil {
			log_colors.CLog('r', "Error closing the database connection:", err)
		} else {
			log_colors.CLog('g', "Database connection closed successfully.")
		}
	}
}
