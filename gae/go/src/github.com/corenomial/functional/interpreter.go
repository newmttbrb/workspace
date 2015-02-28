package main

import (
	"code.google.com/p/go.net/websocket"
	"html/template"
	"io"
	"log"
	"net/http"
)

type command_message struct {
	Target    string `json:"t"`
	Command   string `json:"c"`
	Arguments string `json:"a"`
}

type result_message struct {
	Target    string `json:"t"`
	Command   string `json:"c"`
	Arguments string `json:"a"`
	Result    string `json:"r"`
}

var (
	list          broadcastList
	listenAddress = "localhost:1337"
)

func serve_static_content(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", serve)
	http.Handle("/ws", websocket.Handler(chat))
	http.HandleFunc("/static/", serve_static_content)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	var page = template.
			Must(template.New("index.html").
			Delims("{$", "$}").
			ParseFiles("index.html"))

	page.Execute(w, listenAddress)
}

func chat(c *websocket.Conn) {
	m := &command_message{Arguments: "Hello, welcome to the Tester"}
	websocket.JSON.Send(c, *m)

	list.AppendConnection(c)

	log.Println("starting poling")
	for {
		err := websocket.JSON.Receive(c, m)
		if err != nil {
			if err == io.EOF {
				log.Println("socket closed!!! cleaning up")
				list.RemoveConnection(c)
				c.Close()
				return
			}
			log.Println("error received " + err.Error())
		} else {
			log.Println("received [" + m.Target + ":" + m.Command + "> \"" + m.Arguments + "\"]")
		}

		list.Broadcast(m)
	}
}
