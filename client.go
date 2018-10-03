package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

// var msg string

func main() {
	fmt.Println("Cliner runing...")

	// create ecoonection
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/echo", nil)

	if err != nil {
		log.Println("dail error: ", err)
	}

	run(conn)
}

func run(c *websocket.Conn) {
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please input a data:")

		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		err := c.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("write:", err)
			return
		}

		_, echo, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			return
		}

		log.Println("recv:", string(echo))
	}
}
