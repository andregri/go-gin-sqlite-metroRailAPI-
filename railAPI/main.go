package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/andregri/go-restful-sqlite-metroRailAPI/dbutils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Global variable to hold the DB driver
var DB *sql.DB

// StationResource holds the information about the location
type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

// GetStation returns the station details
func GetStation(c *gin.Context) {
	// Query the station from the DB by id
	var station StationResource
	id := c.Param("station-id")
	err := DB.QueryRow(`
		SELECT id, name, CAST(opening_time AS CHAR), CAST(closing_time AS CHAR)
		FROM station WHERE id=?
	`, id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

// CreateStation handles the POST
func CreateStation(c *gin.Context) {
	// Parse the JSON body into the variable
	var station StationResource

	// Load the data into station
	// This is why the StationResource struct has the json inference strings
	if err := c.BindJSON(&station); err != nil {
		// Error decoding the JSON body
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		// Add the resource to the DB
		statement, err := DB.Prepare(`
			INSERT INTO station (name, opening_time, closing_time)
			VALUES (?, ?, ?)
		`)
		if err != nil {
			log.Println("insert statement prepare error:", err)
		}

		result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err != nil {
			// Return an error because it was not possible to create a row
			log.Println("insert statement execution error:", err)
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			// Successfully created a new row
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		}
	}
}

// RemoveStation handles the DELETE
func RemoveStation(c *gin.Context) {
	// Strip the station-id parameter from the path
	id := c.Param("station-id")

	statement, err := DB.Prepare(`
		DELETE FROM station WHERE id=?
	`)
	if err != nil {
		log.Println("delete stmt prepare error:", err)
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		_, err := statement.Exec(id)
		if err != nil {
			log.Println("delele stmt exec error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			log.Println("Successfully deleted station")
			c.JSON(http.StatusOK, gin.H{
				"result": id,
			})
		}
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("DB driver creation failed!")
	}

	dbutils.Initialize(DB)

	r := gin.Default()

	// Add routes
	r.GET("/v1/stations/:station-id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station-id", RemoveStation)

	r.Run(":8000")
}
