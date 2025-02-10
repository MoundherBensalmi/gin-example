package main

import (
	"MBFacto/app/models"
	"MBFacto/config"
	"MBFacto/database"
	"MBFacto/utils/log_colors"
	"flag"
	"log"
)

var modelsList = []interface{}{
	&models.User{},
}

func init() {
	config.Load()
	database.ConnectToDB()
}

func migrate() {
	db := database.GetDB()
	for _, model := range modelsList {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatal(err)
		}
	}

	log_colors.CLog('g', "Migration completed successfully")
}

func deleteTables() {
	db := database.GetDB()
	for _, model := range modelsList {
		err := db.Migrator().DropTable(model)
		if err != nil {
			log_colors.CLog('r', "Error dropping table for model", model, ":", err)
		}
	}

	log_colors.CLog('y', "Deleted all tables")
}

func main() {
	freshFlag := flag.Bool("f", false, "If set, performs a fresh migration")
	flag.Parse()

	if *freshFlag {
		log_colors.CLog('b', "Performing a fresh migration...")
		deleteTables()
	}
	migrate()
}
