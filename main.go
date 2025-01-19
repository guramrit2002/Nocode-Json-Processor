package main

import (
	"context"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/guramrit2002/Nocode-Json-Processor.git/Db"
	"github.com/guramrit2002/Nocode-Json-Processor.git/jsonHandelers"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {

	// Setting Database
	conn = Db.InitDB(conn)
	Db.CreateTables(conn)

	defer func() {
		if conn != nil {
			conn.Close(context.Background())
		}
	}()

	// Set up Gin router
	r := gin.Default()

	// Health-check endpoint
	r.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Go Server Is Running!"})
	})

	// Saving Widgets endpoint
	r.POST("/data", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)

		if err != nil {
			log.Printf("Error reading request body: %v", err)
			c.JSON(400, gin.H{"status": "error", "message": "Invalid request body"})
			return
		}
		log.Printf("Request Body: %s", body)

		// Simulate WidgetHandler
		var status, message, result = jsonHandelers.CreateWidgetHandler(conn, body)

		// Respond with JSON
		c.JSON(200, gin.H{
			"status":  status,
			"message": message,
			"result":  result,
		})
	})

	// Run the server
	r.Run("0.0.0.0:10000") // Default: localhost:8080
}
