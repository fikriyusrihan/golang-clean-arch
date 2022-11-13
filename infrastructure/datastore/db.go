package datastore

import (
	"fmt"
	"os"

	"github.com/fikriyusrihan/golang-clean-arch/config"
	"github.com/fikriyusrihan/golang-clean-arch/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	DSN := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.DBName,
		config.C.Database.Username,
		config.C.Database.Password,
		config.C.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.AutoMigrate(&domain.Book{}, &domain.Review{})

	fmt.Println("Database instance created")

	return db
}
