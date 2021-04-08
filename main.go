package main

import (
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

	for _, addr := range strings.Split(strings.ReplaceAll(string(b), "\r\n", "\n"), "\n") {
		conn, res, err := websocket.DefaultDialer.Dial(addr, nil)
		if err != nil {
			log.Printf("\033[31m connect to %s, err: %s \033[0m", addr, err)
		} else {
			log.Printf("\033[32m connect to %s, res: %+v \033[0m", addr, res)
			conn.Close()
		}
	}
}
