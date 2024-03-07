package seeders

import (
	"fmt"
	"github.com/MarioMedWilson/car-catalog/models"
	"gorm.io/gorm"
)

func SeedCar(db *gorm.DB){
	fmt.Println("Seeding Cars...")
	db.Create(&models.Car{Name: "C-Class C63 AMG Coupe Aerodynamic Retrofit", Make: "Mercedes benz", ModelYear: 1965, TypeID: 3, Color: "Black", SpeedRange: "0-300km/h"})
	db.Create(&models.Car{Name: "Mustang GT Fastback", Make: "Ford", ModelYear: 1967, TypeID: 1, Color: "Dark Blue", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Camaro SS", Make: "Chevrolet", ModelYear: 1969, TypeID: 4, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Challenger SRT Hellcat", Make: "Dodge", ModelYear: 1970, TypeID: 1, Color: "Black", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Corvette Stingray", Make: "Chevrolet", ModelYear: 1970, TypeID: 1, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "911 Carrera", Make: "Porsche", ModelYear: 1973, TypeID: 1, Color: "Black", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "RX-7", Make: "Mazda", ModelYear: 1978, TypeID: 5, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "300ZX", Make: "Nissan", ModelYear: 1983, TypeID: 1, Color: "Black", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Supra", Make: "Toyota", ModelYear: 1986, TypeID: 6, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Skyline GT-R", Make: "Nissan", ModelYear: 1989, TypeID: 1, Color: "Black", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "NSX", Make: "Honda", ModelYear: 1990, TypeID: 11, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Viper", Make: "Dodge", ModelYear: 1992, TypeID: 12, Color: "Black", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Z3", Make: "BMW", ModelYear: 1996, TypeID: 2, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "S2000", Make: "Honda", ModelYear: 1999, TypeID: 7, Color: "Black", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "RX-8", Make: "Mazda", ModelYear: 2003, TypeID: 9, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "350Z", Make: "Nissan", ModelYear: 2003, TypeID: 5, Color: "Black", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Cayman", Make: "Porsche", ModelYear: 2005, TypeID: 2, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Challenger", Make: "Dodge", ModelYear: 2008, TypeID: 1, Color: "Black", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "370Z", Make: "Nissan", ModelYear: 2009, TypeID: 5, Color: "Red", SpeedRange: "0-220km/h"})
	db.Create(&models.Car{Name: "Genesis Coupe", Make: "Hyundai", ModelYear: 2009, TypeID: 1, Color: "Black", SpeedRange: "0-220km/h"})
}
