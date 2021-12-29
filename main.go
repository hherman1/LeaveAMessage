package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
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

	log.Println("Received Message \"" + bodyStr + "\" from address " + r.Header.Get("X-Real-IP"))

	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "http://static.hherman.com")
	_, err = fmt.Fprintln(w, reply)
	if err != nil {
		log.Println(err)
	}
}

var port = os.Getenv("PORT")

//go:generate esbuild web/src/main.ts --bundle --outfile=web/out.js
//go:embed web
var web embed.FS

func main() {
	http.HandleFunc("/note", handler)
	sub, err := fs.Sub(web, "web")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	http.Handle("/", http.FileServer(http.FS(sub)))

	if port == "" {
		port = "8083"
	}
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
