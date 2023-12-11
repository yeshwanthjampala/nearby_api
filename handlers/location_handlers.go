package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/yeshwanthjampala/nearby_api/database"
	"github.com/yeshwanthjampala/nearby_api/models"
	"github.com/yeshwanthjampala/nearby_api/utils"

	"github.com/gin-gonic/gin"
)

// // JWT Middleware to validate tokens
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")

// 		if tokenString == "" {
// 			utils.JSONResponse(c, http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
// 			c.Abort()
// 			return
// 		}

// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("JWT_SECRET"), nil // Use your actual secret key here
// 		})

// 		if err != nil || !token.Valid {
// 			utils.JSONResponse(c, http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func CreateToken(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")

// 	// Authenticate user (validate username and password)
// 	// Example: Check credentials against your user database

// 	if username == "valid_username" && password == "valid_password" {
// 		token := jwt.New(jwt.SigningMethodHS256)
// 		claims := token.Claims.(jwt.MapClaims)

// 		claims["username"] = username
// 		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

// 		tokenString, _ := token.SignedString([]byte("JWT_SECRET")) // Use your actual secret key here

// 		utils.JSONResponse(c, http.StatusOK, gin.H{"token": tokenString})
// 		return
// 	}

// 	utils.JSONResponse(c, http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// }

// MeasureResponseTime middleware to measure endpoint response time
func MeasureResponseTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now() // Record the start time before processing the request

		// Get existing response data
		existingData := getExistingData(c)

		// Handle the request
		c.Next()

		// Calculate response time after processing the request
		endTime := time.Now()                                // Record the end time after processing the request
		responseTime := endTime.Sub(startTime).Nanoseconds() // Calculate response time

		// Add response time and existing data to the JSON response
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"data":    existingData,
			"time_ns": responseTime,
		})
	}
}

// Function to get existing response data
func getExistingData(c *gin.Context) interface{} {
	// Example 1: Retrieving data from the request context
	existingData, exists := c.Get("existing_data_key")
	if exists {
		return existingData
	}

	// Example 2: Retrieving data from a database or other storage
	// Replace this with your database retrieval logic
	// For instance:
	// dataFromDB := database.GetExistingDataByID(id)
	// return dataFromDB

	// Example 3: Returning static data
	return "Existing data"
}

// func calculateTripCostWithTollGuru(destinationID string, userLocation models.UserLocation, apiKey string) (models.TripCost, error) {
// 	// Validate the destinationID
// 	if destinationID == "" {
// 		return models.TripCost{}, errors.New("destinationID is required")
// 	}

// 	// Validate the userLocation fields
// 	if userLocation.Latitude == 0 || userLocation.Longitude == 0 {
// 		return models.TripCost{}, errors.New("invalid user location")
// 	}

// 	// Validate the apiKey
// 	if apiKey == "" {
// 		return models.TripCost{}, errors.New("TollGuru API key is missing")
// 	}

// 	// Construct the request URL using the TollGuru API endpoint and parameters
// 	apiURL := fmt.Sprintf("https://tollguru.com/v1/calculate?apiKey=%s&source=%f,%f&destination=%s",
// 		apiKey, userLocation.Latitude, userLocation.Longitude, destinationID)

// 	// Make an HTTP GET request to the TollGuru API
// 	resp, err := http.Get(apiURL)
// 	if err != nil {
// 		return models.TripCost{}, fmt.Errorf("failed to fetch data from TollGuru API: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Check for a successful response status
// 	if resp.StatusCode != http.StatusOK {
// 		return models.TripCost{}, fmt.Errorf("TollGuru API returned an error: %s", resp.Status)
// 	}

// 	// Read the API response body
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return models.TripCost{}, fmt.Errorf("failed to read response body: %v", err)
// 	}

// 	// Parse the API response into the TripCost struct
// 	var tripCost models.TripCost
// 	if err = json.Unmarshal(body, &tripCost); err != nil {
// 		return models.TripCost{}, fmt.Errorf("failed to parse TollGuru API response: %v", err)
// 	}

// 	return tripCost, nil
// }

// func GetTripCostHandler(c *gin.Context) {
// 	locationID := c.Param("location_id")

// 	// Extract user's current location from the request body
// 	var userLocation models.UserLocation
// 	if err := c.BindJSON(&userLocation); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format for user location"})
// 		return
// 	}

// 	// Retrieve the TollGuru API key from environment variables
// 	apiKey := os.Getenv("TOLLGURU_API_KEY")
// 	if apiKey == "" {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "TollGuru API key not found"})
// 		return
// 	}

// 	// Call the function to calculate trip cost using the provided API key and user's location
// 	// destinationID := c.Param("location_id")
// 	tripCost, err := calculateTripCostWithTollGuru(locationID, userLocation, apiKey)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Return the calculated trip cost as a JSON response
// 	c.JSON(http.StatusOK, tripCost)
// }

func calculateTripCostWithTollGuru(destinationID string, userLocation models.UserLocation, apiKey string) (models.TripCost, error) {
	// Validate the destinationID
	if destinationID == "" {
		return models.TripCost{}, errors.New("destinationID is required")
	}

	// Validate the userLocation fields
	if userLocation.Latitude == 0 || userLocation.Longitude == 0 {
		return models.TripCost{}, errors.New("invalid user location")
	}

	// Validate the apiKey
	if apiKey == "" {
		return models.TripCost{}, errors.New("TollGuru API key is missing")
	}

	// Construct the request URL using the TollGuru API endpoint and parameters
	apiURL := fmt.Sprintf("https://tollguru.com/v1/calculate?apiKey=%s&source=%f,%f&destination=%s",
		apiKey, userLocation.Latitude, userLocation.Longitude, destinationID)

	// Make an HTTP GET request to the TollGuru API
	resp, err := http.Get(apiURL)
	if err != nil {
		return models.TripCost{}, fmt.Errorf("failed to fetch data from TollGuru API: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response status
	if resp.StatusCode != http.StatusOK {
		return models.TripCost{}, fmt.Errorf("TollGuru API returned an error: %s", resp.Status)
	}

	// Read the API response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.TripCost{}, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the API response into the TripCost struct
	var tripCost models.TripCost
	if err = json.Unmarshal(body, &tripCost); err != nil {
		// Check for non-JSON responses or errors in the API response
		if strings.Contains(string(body), "<") {
			return models.TripCost{}, errors.New("TollGuru API returned a non-JSON response")
		}
		return models.TripCost{}, fmt.Errorf("failed to parse TollGuru API response: %v", err)
	}

	return tripCost, nil
}

func GetTripCostHandler(c *gin.Context) {
	locationID := c.Param("location_id")

	// Extract user's current location from the request body
	var userLocation models.UserLocation
	if err := c.BindJSON(&userLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format for user location"})
		return
	}

	// Retrieve the TollGuru API key from environment variables
	apiKey := os.Getenv("TOLLGURU_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TollGuru API key not found"})
		return
	}

	// Call the function to calculate trip cost using the provided API key and user's location
	tripCost, err := calculateTripCostWithTollGuru(locationID, userLocation, apiKey)
	if err != nil {
		// Check for specific error messages related to API response
		if strings.Contains(err.Error(), "non-JSON response") {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected response from TollGuru API"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the calculated trip cost as a JSON response
	c.JSON(http.StatusOK, tripCost)
}

func GetLocationsByCategorys(c *gin.Context) {
	category := c.Param("category")
	db := database.GetDB() // Assuming database.GetDB() retrieves the database connection

	var locations []models.Location
	result := db.Where("category = ?", category).Find(&locations)
	if result.Error != nil {
		log.Fatal("Error querying database:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}

	response := gin.H{"locations": locations}
	c.JSON(http.StatusOK, response)
}

func SearchLocations(c *gin.Context) {
	var searchParams models.SearchParams

	if err := c.ShouldBindJSON(&searchParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Simulated database query logic based on search parameters
	nearbyLocations, err := GetNearbyLocations(searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch nearby locations"})
		return
	}

	response := gin.H{"locations": nearbyLocations}
	c.JSON(http.StatusOK, response)
}

// Function to fetch nearby locations based on search parameters
func GetNearbyLocations(params models.SearchParams) ([]models.Location, error) {
	var nearbyLocations []models.Location

	// Retrieving the database connection
	db := database.GetDB()
	if db == nil {
		return nil, errors.New("database connection error")
	}

	// // Fetching nearby locations using GORM methods
	// if err := db.Where("category = ? AND ST_DWithin(location, ST_MakePoint(?, ?)::geography, ?)", params.Category, params.Longitude, params.Latitude, params.RadiusKm*1000).Find(&nearbyLocations).Error; err != nil {
	// 	return nil, err
	// }

	// Fetching nearby locations using GORM methods
	if err := db.Where("category = ? AND ST_DWithin(ST_MakePoint(longitude, latitude)::geography, ST_MakePoint(?, ?)::geography, ?)", params.Category, params.Longitude, params.Latitude, params.RadiusKm*1000).Find(&nearbyLocations).Error; err != nil {
		return nil, err
	}

	return nearbyLocations, nil
}

func CreateLocation(c *gin.Context) {
	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add custom validations if needed before creating the location record

	db := database.GetDB()
	if db == nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	if err := db.Create(&location).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Failed to create location"})
		return
	}

	utils.JSONResponse(c, http.StatusCreated, location)
}

func GetLocationByID(c *gin.Context) {
	var location models.Location
	id := c.Param("id")

	db := database.GetDB()
	if db == nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	if err := db.First(&location, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	utils.JSONResponse(c, http.StatusOK, location)
}

func GetAllLocations(c *gin.Context) {
	var locations []models.Location

	db := database.GetDB()
	if db == nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	if err := db.Find(&locations).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch locations"})
		return
	}

	utils.JSONResponse(c, http.StatusOK, locations)
}

func UpdateLocationByID(c *gin.Context) {
	var location models.Location
	id := c.Param("id")

	db := database.GetDB()
	if db == nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	if err := db.First(&location, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	// Update the location based on the JSON data received
	if err := c.ShouldBindJSON(&location); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add custom validations/logic before updating the location record

	if err := db.Save(&location).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Failed to update location"})
		return
	}

	utils.JSONResponse(c, http.StatusOK, location)
}

func DeleteLocationByID(c *gin.Context) {
	var location models.Location
	id := c.Param("id")

	db := database.GetDB()
	if db == nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	if err := db.First(&location, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	// Add any additional checks or validations before deleting the location

	if err := db.Delete(&location).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Failed to delete location"})
		return
	}

	utils.JSONResponse(c, http.StatusOK, gin.H{"message": "Location deleted successfully"})
}
