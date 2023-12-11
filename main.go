package main

import (
	"log"
	"github.com/yeshwanthjampala/nearby_api/database"
	"github.com/yeshwanthjampala/nearby_api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	db := database.InitDB()
	defer db.Close()

	// Apply MeasureResponseTime middleware to all routes
	r.Use(handlers.MeasureResponseTime())

	// Routes
	//r.POST("/token", handlers.CreateToken)
	// r.GET("/locations/:category", handlers.AuthMiddleware(), handlers.GetLocationsByCategorys)
	// r.POST("/search", handlers.AuthMiddleware(), handlers.SearchLocations)
	// r.POST("/locations", handlers.AuthMiddleware(), handlers.CreateLocation)
	// r.GET("/location/:id", handlers.AuthMiddleware(), handlers.GetLocationByID)
	// r.GET("/locations", handlers.AuthMiddleware(), handlers.GetAllLocations)
	// r.PUT("/locations/:id", handlers.AuthMiddleware(), handlers.UpdateLocationByID)
	// r.DELETE("/locations/:id", handlers.AuthMiddleware(), handlers.DeleteLocationByID)
	// r.POST("/trip-cost/:location_id", handlers.AuthMiddleware(), handlers.GetTripCostHandler)

	r.GET("/locations/:category", handlers.GetLocationsByCategorys)
	r.POST("/search", handlers.SearchLocations)
	r.POST("/locations", handlers.CreateLocation)
	r.GET("/location/:id", handlers.GetLocationByID)
	r.GET("/locations", handlers.GetAllLocations)
	r.PUT("/locations/:id", handlers.UpdateLocationByID)
	r.DELETE("/locations/:id", handlers.DeleteLocationByID)
	r.POST("/trip-cost/:location_id", handlers.GetTripCostHandler)

	log.Fatal(r.Run(":8080"))
}
