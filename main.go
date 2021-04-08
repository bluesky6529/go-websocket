package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	var b, str string
	for i := 1; i < len(os.Args); i++ {
		b += str + os.Args[i]
		str = "\n"
	}
	fmt.Println(b)

	for _, addr := range strings.Split(strings.ReplaceAll(string(b), "\r\n", "\n"), "\n") {
		conn, res, err := websocket.DefaultDialer.Dial(addr, nil)
		if err != nil {
			log.Printf("connect to %s, err: %s", addr, err)
		} else {
			log.Printf("connect to %s, res: %+v", addr, res)
			conn.Close()
		}
	}
	//log.Printf("Press the any key to terminate.")
	//fmt.Scanln(&b)
}
