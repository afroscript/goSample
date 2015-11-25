package main

import (
	"code.google.com/p/go.net/websocket"
	//"io"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"
	"os"
)

func echoHandler(ws *websocket.Conn) {
	type T struct {
		X     int64
		Y     int64
		Color string
		Str   string
	}

	var data T
	websocket.JSON.Receive(ws, &data)

	log.Printf("%#v\n", data)

	websocket.JSON.Send(ws, data)
}

func initRedis() redis.Conn {
	var h = flag.String("hostname", "127.0.0.1", "Set Hostname")
	var p = flag.String("port", "6379", "Set Port")
	// var key = flag.String("key", "foo", "Set Key")
	// var val = flag.String("value", "bar", "Set Value")
	flag.Parse()

	c, err := redis.Dial("tcp", *h+":"+*p)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return c
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
