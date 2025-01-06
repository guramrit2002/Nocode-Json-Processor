package main

import (
	"fmt"
	"github.com/guramrit2002/Nocode-Json-Processor.git/jsonHandelers"
)

func main() {
	fmt.Println("App is working !")
	var status, Message, Result = jsonHandelers.WidgetHandler()
	fmt.Println(status)
	fmt.Println(Message)
	fmt.Println(Result)
}
