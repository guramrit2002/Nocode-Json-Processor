package Db

import (
	"context"
	"encoding/json"
	"github.com/guramrit2002/Nocode-Json-Processor.git/constants"
	"github.com/guramrit2002/Nocode-Json-Processor.git/models"
	"github.com/jackc/pgx/v5"
	"log"
)

// InsertWidget inserts a widget and its children into the database
func InsertWidget(Conn *pgx.Conn, widget models.Widget, parentID *string) error {
	// Convert `Style` to JSON
	styleJSON, err := json.Marshal(widget.Style)
	if err != nil {
		return err
	}

	// SQL query to insert the widget
	query := constants.WidgetInsertCommand

	// Execute the query
	_, err = Conn.Exec(context.Background(), query, widget.ID, widget.Type, styleJSON, widget.Content, parentID)
	if err != nil {
		log.Printf("Failed to insert widget %s: %v\n", widget.ID, err)
		return err
	}

	log.Printf("Widget %s inserted successfully.", widget.ID)

	// Recursively insert children
	for _, child := range widget.Children {
		err := InsertWidget(Conn, child, &widget.ID)
		if err != nil {
			log.Printf("Failed to insert child widget %s: %v\n", child.ID, err)
		}
	}

	return nil
}
