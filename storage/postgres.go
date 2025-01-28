package storage

import (
	"fmt"
	"github.com/OlegacyGold/test/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	connStr := fmt.Sprintf("host=db port=5432 user=user password=password dbname=testdb sslmode=disable")
	for {
		var err error
		DB, err = gorm.Open("postgres", connStr)
		if err == nil {
			break
		}
		log.Println("Postgres is trying to connect...")
		time.Sleep(5 * time.Second)
	}
	if err := DB.AutoMigrate(&models.User{}).Error; err != nil {
		log.Fatal(err)
	}
	if _, err := DB.DB().Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\""); err != nil {
		log.Fatal(err)
	}
	os.Setenv("PGCLIENTENCODING", "utf-8")
}
