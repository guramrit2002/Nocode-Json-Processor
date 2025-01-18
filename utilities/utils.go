package utilities

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"

	"github.com/guramrit2002/Nocode-Json-Processor.git/Db"
	"github.com/guramrit2002/Nocode-Json-Processor.git/models"
)

func iterativeSearchWidget(Conn *pgx.Conn, widgets []models.Widget) []string {
	//copy of the root nodes to avoid modifying the original slice
	stack := append([]models.Widget(nil), widgets...)

	// visited widgets
	var result []string

	for len(stack) > 0 {
		//Popping out the node
		widget := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// process the widget
		// add them to db
		err := Db.InsertWidget(Conn, widget, nil) // Root widgets have no parent (parentID = nil)
		if err != nil {
			log.Printf("Error inserting widget %s: %v", widget.ID, err)
		}
		// Add widget to result slice (store the result)
		result = append(result, widget.ID)

		// appending children to stack
		for i := len(widget.Children) - 1; i >= 0; i-- {
			stack = append(stack, widget.Children[i])
		}
	}

	return result
}

func WidgetJson(Conn *pgx.Conn, body []byte) (bool, string, []string) {
	var widgetMap []models.Widget
	err := json.Unmarshal(body, &widgetMap)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return false, "Failed", []string{}
	}
	var result = iterativeSearchWidget(Conn, widgetMap)
	if len(result) > 0 {
		return true, "success", result
	}
	return true, "Success", result
}
