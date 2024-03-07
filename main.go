package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/MarioMedWilson/car-catalog/routes"
	"github.com/MarioMedWilson/car-catalog/config"
	"github.com/MarioMedWilson/car-catalog/migrations"
	"github.com/MarioMedWilson/car-catalog/seeders"
)


func main() {
	db, err := config.GetDB()
	if err != nil {
		panic(err)
	}
	migrations.Migrate(db)

	seeders.SeedCarTypes(db)
	seeders.SeedCar(db)

	router := gin.Default()
	routes.SetupRouter(router, db)
	
	router.Run(":9090")
	fmt.Println("Server running on port 9090")
}
