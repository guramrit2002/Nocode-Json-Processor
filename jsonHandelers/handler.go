package jsonHandelers

import (
	"github.com/guramrit2002/Nocode-Json-Processor.git/utilities"
	"github.com/jackc/pgx/v5"
)

func CreateWidgetHandler(Conn *pgx.Conn, body []byte) (bool, string, []string) {
	var status, Message, Result = utilities.WidgetJson(Conn, body)
	if !status {
		return false, Message, Result
	}
	return true, Message, Result
}
