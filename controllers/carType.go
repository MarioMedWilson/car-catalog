package controllers

import (
	"github.com/MarioMedWilson/car-catalog/models"
	// "fmt"
	"gorm.io/gorm"
)


type ReturnMessage struct {
	Message string `json:"message"`
	Status int 
}

func GetCarTypes(db *gorm.DB) []models.CarType {
	var carTypes []models.CarType
	db.Find(&carTypes)
	return carTypes
}

func CreateCarType(db *gorm.DB, carType models.CarType) ReturnMessage {
	if carType.Type == "" {
		return ReturnMessage{Message: "Type is required", Status: 400}
	}
	db.Create(&carType)
	return ReturnMessage{Message: "Car Type created successfully", Status: 200}
}

func GetCarTypeByID(db *gorm.DB, id string) (models.CarType, ReturnMessage) {
	var carType models.CarType
	db.First(&carType, id)
	if carType.ID == 0 {
		return carType, ReturnMessage{Message: "Car Type not found", Status: 404}
	}
	return carType, ReturnMessage{Message: "Car Type found", Status: 200}
}

func UpdateCarType(db *gorm.DB, id string, carType models.CarType) ReturnMessage {
	var carTypeToUpdate models.CarType
	db.First(&carTypeToUpdate, id)
	if carTypeToUpdate.ID == 0 {
		return ReturnMessage{Message: "Car Type not found", Status: 404}
	}
	db.Model(&carTypeToUpdate).Updates(carType)
	return ReturnMessage{Message: "Car Type updated successfully", Status: 200}
}

func DeleteCarType(db *gorm.DB, id string) ReturnMessage {
	var carType models.CarType
	db.First(&carType, id)
	if carType.ID == 0 {
		return ReturnMessage{Message: "Car Type not found", Status: 404}
	}
	db.Delete(&carType)
	return ReturnMessage{Message: "Car Type deleted successfully", Status: 200}
}
