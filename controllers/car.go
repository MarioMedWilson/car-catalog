package controllers

import (
	"gorm.io/gorm"
	"github.com/MarioMedWilson/car-catalog/models"
	"encoding/json"
	"strings"
)


func GetCars(db *gorm.DB) []models.Car {
	var cars []models.Car
	db.Preload("Type").Find(&cars)
	return cars
}

func CreateCar(db *gorm.DB, car models.Car) ReturnMessage {
	if car.Name == "" {
		return ReturnMessage{Message: "Model Name is required", Status: 400}
	}
	if car.TypeID == 0 {
		return ReturnMessage{Message: "Type is required", Status: 400}
	}
	if car.Make == "" {
		return ReturnMessage{Message: "Make is required", Status: 400}
	}
	if car.ModelYear == 0 {
		return ReturnMessage{Message: "Model Year is required", Status: 400}
	}
	if car.Color == "" {
		return ReturnMessage{Message: "Color is required", Status: 400}
	}
	if car.SpeedRange == "" {
		return ReturnMessage{Message: "Speed Range is required", Status: 400}
	}

	result := db.Create(&car)

	if result.Error != nil {
		return ReturnMessage{Message: result.Error.Error(), Status: 400}
	}

	return ReturnMessage{Message: "Car created successfully", Status: 200}
}


func GetCarByID(db *gorm.DB, id string) (models.Car, ReturnMessage) {
	var car models.Car
	db.Preload("Type").First(&car, id)
	if car.ID == 0 {
		return car, ReturnMessage{Message: "Car not found", Status: 404}
	}
	return car, ReturnMessage{Message: "Car found", Status: 200}
}

func UpdateCar(db *gorm.DB, id string, car models.Car) ReturnMessage {
	var carToUpdate models.Car
	if car == (models.Car{}) {
		return ReturnMessage{Message: "Car is required", Status: 400}
	}
	db.First(&carToUpdate, id)
	if carToUpdate.ID == 0 {
		return ReturnMessage{Message: "Car not found", Status: 404}
	}
	db.Model(&carToUpdate).Updates(car)
	return ReturnMessage{Message: "Car updated successfully", Status: 200}
}

func DeleteCar(db *gorm.DB, id string) ReturnMessage {
	var car models.Car
	db.First(&car, id)
	if car.ID == 0 {
		return ReturnMessage{Message: "Car not found", Status: 404}
	}
	db.Delete(&car)
	return ReturnMessage{Message: "Car deleted successfully", Status: 200}
}


var FilterData struct {
	Name string `json:"name"`
	Make string `json:"make"`
	ModelYear int `json:"modelYear"`
	Type string `json:"type"`
	Color string `json:"color"`
	SpeedRange string `json:"speedRange"`
}

func GetCarsByFilter(db *gorm.DB, filter string) ([]interface{}, error) {
	var cars []models.Car
	filter = strings.ReplaceAll(filter, "'", "\"")
	err := json.Unmarshal([]byte(filter), &FilterData)
	if err != nil {
		return nil, err
	}

	var results []interface{}
	db.Preload("Type").Find(&cars)
	for _, car := range cars {
		if FilterData.Name != "" && strings.ToLower(car.Name) != strings.ToLower(FilterData.Name) {
			continue
		}
		if FilterData.Make != "" && strings.ToLower(car.Make) != strings.ToLower(FilterData.Make) {
			continue
		}
		if FilterData.ModelYear != 0 && car.ModelYear != FilterData.ModelYear {
			continue
		}
		if FilterData.Type != "" && strings.ToLower(car.Type.Type) != strings.ToLower(FilterData.Type) {
			continue
		}
		if FilterData.Color != "" && strings.ToLower(car.Color) != strings.ToLower(FilterData.Color) {
			continue
		}
		if FilterData.SpeedRange != "" && strings.ToLower(car.SpeedRange) != strings.ToLower(FilterData.SpeedRange) {
			continue
		}

		result := map[string]interface{}{
			"ID":         car.ID,
			"Name":       car.Name,
			"Make":       car.Make,
			"ModelYear":  car.ModelYear,
			"Type":       car.Type.Type,
			"Color":      car.Color,
			"SpeedRange": car.SpeedRange,
		}
		results = append(results, result)
	}

	return results, nil
}
