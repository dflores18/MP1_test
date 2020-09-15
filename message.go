package messages

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Message struct {
	message string
}

func (m *Message) String() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Message: ")
	text, _ := reader.ReadString('\n')
	text = m.message
	return text
}

func messageFrom(message string /* config id */) {
	fmt.Println("\n---------------------------- Incoming Message --------------------------------")
	fmt.Println("Message From: ") // enter identifier
	fmt.Println("Message Content: " + message)
	t := time.Now()
	thisTime := t.Format(time.RFC1123) + "\n"
	fmt.Println("Received At: " + thisTime)
	fmt.Println("\n -----------------------------------------------------------------------------")
}

func messageTo(message string /*config id*/) {
	t := time.Now()
	thisTime := t.Format(time.RFC1123) + "\n"
	fmt.Println("\n---------------------------- Incoming Message --------------------------------")
	fmt.Println("Sent " + "'" + message + "'" + " to process" + /* id */ " " + "system time is " + thisTime) // enter identifier
	fmt.Println("\n -----------------------------------------------------------------------------")
}
