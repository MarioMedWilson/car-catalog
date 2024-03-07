package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/MarioMedWilson/car-catalog/models"
)

func Migrate(db *gorm.DB) {
	fmt.Println("Dropping existing tables...")
	dropTables(db, &models.Car{})
	dropTables(db, &models.CarType{})
	db.AutoMigrate(&models.Car{}, &models.CarType{})
	fmt.Println("Migration complete")
}


func dropTables(db *gorm.DB, model interface{}) {
	if db.Migrator().HasTable(model) {
		db.Migrator().DropTable(model)
	}
}