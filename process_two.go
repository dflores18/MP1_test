package processes

import (
	"../config_file"
	"fmt"
	"io/ioutil"
	"net"
)

type ProcessTwo struct {
	config config_file.Config
}

func unicast_receive(destination string, message string) {
	conn, err := net.Dial("tcp", "localhost:9000") //dials/connects to the same server (tcp)
	if err != nil {
		panic(err)
	} //handles error
	defer conn.Close() //closes connection

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}
