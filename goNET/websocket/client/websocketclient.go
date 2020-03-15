package main

import (
	"fmt"
	// "net/http"
	"bufio"
	"os"
	"time"

	websocket "golang.org/x/net/websocket"
)

func main() {
	ws, err := websocket.Dial("ws://localhost:12345/websocket", "",
		"http://localhost/")
	if err != nil {
		panic("Dial: " + err.Error())
	}
	go readFromServer(ws)
	go writeFromServer(ws)
	time.Sleep(5e9)
	ws.Close()
}

func readFromServer(ws *websocket.Conn) {
	buf := make([]byte, 1000)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}
	}
}

func writeFromServer(ws *websocket.Conn) {
	defer ws.Close()
	file, _ := os.Open("F:\\GoProject\\src\\github.com\\zhangzt123\\Golearn\\goNET\\websocket\\client\\websocketclient.go")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		bytes, _ := reader.ReadBytes(10)
		if _, err := ws.Write(bytes); err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}
	}
}
