package main

import (
	"log"
	"sync"

	"code.google.com/p/go.net/websocket"
)

type broadcastList struct {
	sync.Mutex
	Connections []*websocket.Conn
}

func (l *broadcastList) AppendConnection(c *websocket.Conn) {
	l.Lock()
	defer l.Unlock()

	l.Connections = append(l.Connections, c)
}

func (l *broadcastList) RemoveConnection(c *websocket.Conn) {
	l.Lock()
	defer l.Unlock()

	var newConnections []*websocket.Conn
	for _, cin := range l.Connections {
		if c != cin {
			newConnections = append(newConnections, cin)
		}
	}
	l.Connections = newConnections
}

func (l *broadcastList) Broadcast(m *command_message) {
	l.Lock()
	defer l.Unlock()

	result := &result_message{m.Target, m.Command, m.Arguments, "success"}
	log.Println("sending [" + result.Target + ":" + result.Command + "> \"" + result.Arguments + "\"] result=" + result.Result)
	for _, c := range l.Connections {
		err := websocket.JSON.Send(c, *result)
		if err != nil {
			log.Println(err)
		}
	}
}
