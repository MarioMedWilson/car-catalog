package seeders

import (
	"fmt"
	"github.com/MarioMedWilson/car-catalog/models"
	"gorm.io/gorm"
)

func SeedCarTypes(db *gorm.DB) {
	fmt.Println("Seeding Car Types...")
	db.Create(&models.CarType{Type: "Cabriolet"})
	db.Create(&models.CarType{Type: "Coupe"})
	db.Create(&models.CarType{Type: "Sedan"})
	db.Create(&models.CarType{Type: "Saloon"})
	db.Create(&models.CarType{Type: "Hatchback"})
	db.Create(&models.CarType{Type: "SUV"})
	db.Create(&models.CarType{Type: "MPV"})
	db.Create(&models.CarType{Type: "Crossover"})
	db.Create(&models.CarType{Type: "Convertible"})
	db.Create(&models.CarType{Type: "Wagon"})
	db.Create(&models.CarType{Type: "Van"})
	db.Create(&models.CarType{Type: "Roadstar"})
	db.Create(&models.CarType{Type: "Truck"})
	db.Create(&models.CarType{Type: "Bus"})
	db.Create(&models.CarType{Type: "Supercar"})
	db.Create(&models.CarType{Type: "Mini-van"})
	db.Create(&models.CarType{Type: "Hybrid"})
	db.Create(&models.CarType{Type: "Electric"})
	db.Create(&models.CarType{Type: "Motor-bike"})
	db.Create(&models.CarType{Type: "Four-by-Four"})
	fmt.Println("Car Types seeded successfully")
}
