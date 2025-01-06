package utilities

import (
	"encoding/json"
	"fmt"

	"github.com/guramrit2002/Nocode-Json-Processor.git/data"
	"github.com/guramrit2002/Nocode-Json-Processor.git/structures"
)

func iterativeSearchWidget(widgets []structures.Widget) []string {
	//copy of the root nodes to avoid modifying the original slice
	stack := append([]structures.Widget(nil), widgets...)

	// visited widgets
	var result []string

	for len(stack) > 0 {
		//Popping out the node
		widget := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// process the widget
		// add them to db
		// Add widget to result slice (store the result)
		result = append(result, widget.ID)

		// appending children to stack
		for i := len(widget.Children) - 1; i >= 0; i-- {
			stack = append(stack, widget.Children[i])
		}
	}

	return result
}

func WidgetJson() (bool, string, []string) {
	var defaultMap []structures.Widget
	err := json.Unmarshal([]byte(data.DefaultJson), &defaultMap)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return false, "Failed", []string{}
	}
	var result = iterativeSearchWidget(defaultMap)
	if len(result) > 0 {
		return true, "success", result
	}
	return true, "Success", result
}
