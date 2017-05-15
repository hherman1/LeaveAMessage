package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var lastMessage struct {
	L       sync.Mutex
	message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, 2048)
	bLen, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		log.Println(err)
		fmt.Fprintln(w, "There was an error.")
		return
	}
	body = body[:bLen]
	bodyStr := string(body)

	lastMessage.L.Lock()
	reply := lastMessage.message
	lastMessage.message = bodyStr
	lastMessage.L.Unlock()

	log.Println("Received Message \"" + bodyStr + "\" from address " + r.RemoteAddr)

	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "http://static.hherman.com")
	_, err = fmt.Fprintln(w, reply)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":"+os.Args[1], nil))
}
