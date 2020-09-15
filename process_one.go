package processes

import (
	"../config_file"
	"../messages"
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type ProcessOne struct {
	config config_file.Config
	m      messages.Message
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(netData)
		if temp == "STOP" {
			break
		}

		result := strconv.Itoa(random()) + "\n"
		c.Write([]byte(result))
	}
	c.Close()
}

func unicast_send(destination string, message *messages.Message) {
	l, err := net.Listen("tcp4", ":"+destination)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
