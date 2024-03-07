package	routes

import (
	"github.com/MarioMedWilson/car-catalog/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/MarioMedWilson/car-catalog/models"
	// "fmt"
	// "reflect"
)

func HandleFetchCars(c *gin.Context, db *gorm.DB) {
	filter := c.Query("filter")
	if filter != "" {
		result, err := controllers.GetCarsByFilter(db, filter)
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
		c.JSON(200, result)
		return
	}

	cars := controllers.GetCars(db)
	c.JSON(200, cars)
}

func SetupCarRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		HandleFetchCars(c, db)
	})

	router.POST("/", func(c *gin.Context) {
		var car models.Car
		c.BindJSON(&car)
		val := controllers.CreateCar(db, car)
		if val.Status == 200 {
			c.JSON(val.Status, gin.H{"message": val.Message})
		} else {
			c.JSON(val.Status, gin.H{"message": val.Message})
		}
	})

	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		car, val := controllers.GetCarByID(db, id)
		if val.Status == 200 {
			c.JSON(val.Status, car)
		} else {
			c.JSON(val.Status, gin.H{"message": val.Message})
		}
	})

	router.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var car models.Car
		c.BindJSON(&car)
		val := controllers.UpdateCar(db, id, car)
		if val.Status == 200 {
			c.JSON(val.Status, gin.H{"message": val.Message})
		} else {
			c.JSON(val.Status, gin.H{"message": val.Message})
		}
	})

	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		val := controllers.DeleteCar(db, id)
		if val.Status == 200 {
			c.JSON(val.Status, gin.H{"message": val.Message})
		} else {
			c.JSON(val.Status, gin.H{"message": val.Message})
		}
	})

}

func SetupCarTypeRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		carTypes := controllers.GetCarTypes(db)
		c.JSON(200, carTypes)
	})

	router.POST("/", func(c *gin.Context) {
		var carType models.CarType
		c.BindJSON(&carType)
		val := controllers.CreateCarType(db, carType)
		if val.Status == 200{
			c.JSON(val.Status, gin.H{"message": val.Message})
		} else {
			c.JSON(val.Status, gin.H{"message": val.Message})
		}
	})

	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		carType, val := controllers.GetCarTypeByID(db, id)
		if val.Status == 200 {
			c.JSON(val.Status, carType)
		} else {
			c.JSON(val.Status, gin.H{"message": val.Message})
		}
	})

	router.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var carType models.CarType
		c.BindJSON(&carType)
		val := controllers.UpdateCarType(db, id, carType)
		if val.Status == 200 {
			c.JSON(val.Status, gin.H{"message": val.Message})
		} else {
			c.JSON(val.Status, gin.H{"message": val.Message})
		}
	})

	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		val := controllers.DeleteCarType(db, id)
		if val.Status == 200 {
			c.JSON(val.Status, gin.H{"message": val.Message})
		} else {
			c.JSON(val.Status, gin.H{"message": val.Message})
		}
	})

}

func SetupRouter(router *gin.Engine, db *gorm.DB) {
	router.GET("/test", controllers.Test)
	carType := router.Group("/car-type")
	{
		SetupCarTypeRoutes(carType, db)
	}
	car := router.Group("/car")
	{
		SetupCarRoutes(car, db)
	}
	router.NoRoute(controllers.NotFound)
}