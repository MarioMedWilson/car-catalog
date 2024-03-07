package config


import (
	"fmt"
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	"strings"
	"github.com/joho/godotenv"
	"os"
	"log"
)

var db *gorm.DB
var err error

func init() {
	err := godotenv.Load()
	fmt.Println(err)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
 
	if err != nil {
		fmt.Println(isDatabaseNotExistError(err))
		if isDatabaseNotExistError(err) {
			fmt.Println("Database does not exist")
			dsn = "host=localhost user=postgres port=5432"
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("Failed to connect to the database")
			}
			db.Exec("CREATE DATABASE carcatalog")
			inctance, _ :=db.DB()
			inctance.Close()
			fmt.Println("Database created")
			dsn = "host=localhost user=postgres dbname=carcatalog port=5432"
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("Failed to connect to the database")
			}
		} else {
			panic("Failed to connect to the database")
		}
	}

	fmt.Println("Connected to the database")
}

func isDatabaseNotExistError(err error) bool {
	if err != nil {
		errStr := err.Error() 
		return strings.Contains(errStr, "3D000")
	}
	return false
}

func GetDB() (*gorm.DB, error){
	return db, err
}
