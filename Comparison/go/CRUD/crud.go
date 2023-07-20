package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Menu represents a model for menus
type Menu struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

func main() {

	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:connect@123@tcp(127.0.0.1:3306)/OES")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Ping the database to ensure the connection is established
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	// Create a Gin router
	r := gin.Default()

	// Read operation - GET /menus
	r.GET("/", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, url,description FROM Menu")
		if err != nil {
			log.Println("Failed to retrieve menus from the database:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve menus"})
			return
		}
		defer rows.Close()

		var menus []Menu
		for rows.Next() {
			var menu Menu
			err := rows.Scan(&menu.ID, &menu.Name, &menu.URL, &menu.Description)
			if err != nil {
				log.Println("Failed to scan menu row:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve menus"})
				return
			}
			menus = append(menus, menu)
		}

		c.JSON(http.StatusOK, menus)
	})

	// Run the server
	r.Run(":9000")
}
