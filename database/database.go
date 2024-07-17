package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Honzikoi/gym-challenge/fixtures"
	"github.com/Honzikoi/gym-challenge/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	var dsn string
	if os.Getenv("DATABASE_URL") != "" {
		// Use Heroku DATABASE_URL
		dsn = os.Getenv("DATABASE_URL")
	} else {
		// Use local environment variables
		dsn = fmt.Sprintf(
			"host=db user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database!")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations...")
	err = db.AutoMigrate(&models.Users{},
		&models.Group{},
		&models.Sessions{},
		&models.Role{},
		&models.About{},
		&models.LoginRequest{},
		&models.LoginResponse{},
		&models.Gym{},
		&models.Equipment{},
		&models.Workout{},
		&models.About{},
		&models.Equipment{},
		&models.Payment{},
		&models.PhysicalBuild{},
		&models.PhysicalObjective{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations! \n", err)
		os.Exit(2)
	}

	DB = Dbinstance{
		Db: db,
	}

	fixtures.FixtureLauncher()
}
