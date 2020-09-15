package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var function string
	var process int
	var message string

	fmt.Print("Please enter a command: ")
	fmt.Scanf("%s %d %s", &function, &process, &message)

	if strings.ToLower(function) == "send" {
		fmt.Print("Sending " + message + " to process" + strconv.Itoa(process))
		//unicast send
		//unicast receive
	} else {
		fmt.Print("invalid command. try 'send'")
	}
}
