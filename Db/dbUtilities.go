package Db

import (
	"github.com/guramrit2002/Nocode-Json-Processor.git/constants"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"context"
	"log"
	"os"
)

func InitDB(conn *pgx.Conn) *pgx.Conn {
	// Connect to the database
	connStr := os.Getenv("DATABASE_URL")
	var errConn error
	conn, errConn = pgx.Connect(context.Background(), connStr)
	if errConn != nil {
		log.Fatalf("Failed to connect to the database: %v", errConn)
	}

	return conn
}

func CreateTables(conn *pgx.Conn) bool {
	// Define the SQL query to create the widgets table
	createTableQuery := constants.WidgetSchemaCommand

	// Log the table creation
	log.Printf("Creating table with query: ", createTableQuery)

	// Execute the query
	_, err := conn.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Printf("Error creating table: %v\n", err)
		return false
	}

	log.Println("Widgets table created successfully!")
	return true
}
