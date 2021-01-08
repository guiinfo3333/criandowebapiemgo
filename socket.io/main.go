package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello World")
	server := socketio.NewServer(nil)

	server.OnConnect("connection", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/socket.io/", server)

	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
