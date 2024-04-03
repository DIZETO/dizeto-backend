package config

import (
	model_about "dizeto-backend/app/model/about"
	model_highlight "dizeto-backend/app/model/highlight_porto"
	model_pricing "dizeto-backend/app/model/pricing"
	model_testimoni "dizeto-backend/app/model/testimoni"
	model_user "dizeto-backend/app/model/user"
	"dizeto-backend/utils"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model_user.User{},
		&model_about.About{},
		&model_highlight.HighlightPortofolio{},
		&model_pricing.Pricing{},
		&model_testimoni.Testimoni{},
	).Error

	if err != nil {
		return nil, err
	}

	// database seeding
	var users = []model_user.User{}
	db.Where("role = ?", "admin").Find(&users)
	fmt.Println(len(users))
	if len(users) == 0 {
		err = SeedUsers(db)
		if err != nil {
			return nil, err
		}

	}

	return db, nil
}

func SeedUsers(db *gorm.DB) error {
	userID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword("admin")
	if err != nil {
		return err
	}
	userAdmin := model_user.User{ID: userID, Username: "admin", Password: hashedPassword, FirstName: "Admin", LastName: "Dizeto", Email: "admin@gmail.com", Role: "admin"}
	db.Create(&userAdmin)

	return nil
}
