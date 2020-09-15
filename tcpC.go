package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	//thisTime := t.Format(time.RFC1123)
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		//sent receipt
		fmt.Println("Sent "+text+"to process "+ /* id */ ", system time is ", &t)
		if strings.TrimSpace(text) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
	//timeout := rand.Intn(3)
	//time.Sleep(time.Duration(timeout) * time.Millisecond)
	//timeout := rand.Intn(5)
	//	c := time.Tick( time.Duration(timeout) * time.Second)
	//	for next := range c {
	//		fmt.Printf("%v %s\n", next, statusUpdate())
	//tried to figure out how to delay the message
	//ended up not using this
}
