package main

import (
	"net"
	"os"
	"strconv"
)

type TCPTurret struct {
	Ip         string
	Port       int
	connection *net.TCPConn
}

func TurretFactory(ip string, port int) (tcpt *TCPTurret) {
	tcpt = new(TCPTurret)
	tcpt.Ip = ip
	tcpt.Port = port
	return tcpt
}

func (tcpt *TCPTurret) Connect() (err error) {

	if tcpt.connection == nil {

		var turretAddress string

		turretAddress = tcpt.Ip + ":" + strconv.Itoa(tcpt.Port)

		tcpAddr, err := net.ResolveTCPAddr("tcp", turretAddress)

		if err == nil {
			tcpt.connection, err = net.DialTCP("tcp", nil, tcpAddr)
		}
	}
	return
}

func (tcpt *TCPTurret) Send(text string) (err error) {
	_, err = tcpt.connection.Write([]byte(text))
	return
}

func (tcpt *TCPTurret) Receive() (str string, err error) {
	reply := make([]byte, 10240)

	_, err = tcpt.connection.Read(reply)
	str = string(reply)
	return
}

func (tcpt *TCPTurret) Disconnect() {
	tcpt.connection.Close()
}

func main() {
	turret := TurretFactory("localhost", 6666)
	err := turret.Connect()

	if err != nil {
		println("could not connect to turret", err.Error())
		os.Exit(1)
	}

	defer turret.Disconnect()

	err = turret.Send("<xmlSample><a>1</a><b>2</b><?xmlSample>")
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	reply, err := turret.Receive()
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

}
