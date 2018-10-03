package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("Server Starting...")

	flag.Parse()

	setRoutes()

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func setRoutes() {
	// setting websocket
	http.HandleFunc("/echo", echo)
	// setting index.html 靜態頁面
	http.HandleFunc("/", index)

}

func echo(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read: %s", err)
			break
		}

		log.Println("recv:", string(message))

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	httpTemplate := template.Must(template.ParseFiles("index.html"))
	httpTemplate.Execute(w, "ws://"+r.Host+"/echo")
}
