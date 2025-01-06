package jsonHandelers

import (
	"github.com/guramrit2002/Nocode-Json-Processor.git/utilities"
)

func WidgetHandler() (bool, string, []string) {
	var status, Message, Result = utilities.WidgetJson()
	if !status {
		return false, Message, Result
	}
	return true, Message, Result
}
