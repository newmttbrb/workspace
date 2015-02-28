// tcp_connection_to_device
package tcp_connection_to_device

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

type TCPConnection struct {
	server     string
	port       int
	name       string
	connection net.Conn
}

func ConnectionFactory(server_arg string, port int, name_arg string) (pconn *TCPConnection) {
	return &TCPConnection{server_arg, port, name_arg, nil}
}

func (pconn *TCPConnection) Start() error {
	if pconn.connection != nil {
		log.Println("(", pconn.name, ") called start TCP without stopping previous connection, stopping")
		pconn.Stop()
	}
	log.Println("(", pconn.name, ") about to Dial TCP")
	c, err := net.Dial("tcp", pconn.server+":"+strconv.Itoa(pconn.port))
	if err != nil {
		log.Println("(", pconn.name, ") could not start the TCP connection to the device ", pconn.name, ", the error was ", err)
		return err
	}
	pconn.connection = c
	log.Println("(", pconn.name, ") successfully started a TCP connection to the device ", pconn.name)
	return nil
}

func (pconn *TCPConnection) Stop() (err error) {
	if pconn.connection != nil {
		err = pconn.connection.Close()
		if err != nil {
			log.Println("(", pconn.name, ") could not stop the TCP connection to the device ", pconn.name, ", the error was ", err)
		} else {
			log.Println("(", pconn.name, ") successfully stopped a TCP connection to the device ", pconn.name)
		}
		pconn.connection = nil
	}
	return
}

func (pconn *TCPConnection) SendCommand(cmd string, use_stdio bool) (err error) {
	_, err = fmt.Fprintf(pconn.connection, cmd)
	if err != nil {
		log.Println("(", pconn.name, ") could not run the command '", cmd, "', the error was ", err)
		return
	}
	reader := bufio.NewReader(pconn.connection)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if use_stdio {
			fmt.Println(string(line))
		}
	}
	log.Println("(", pconn.name, ") ran the command '", cmd, "' on the device ", pconn.name, " without error")
	return
}

func (pconn *TCPConnection) SendCommands(cmds []string, use_stdio bool) (err error) {
	for _, cmd := range cmds {
		err = pconn.SendCommand(cmd, use_stdio)
		if err != nil {
			// no logging here because the SendCommand should have logged it already
			break
		}
	}
	return
}
