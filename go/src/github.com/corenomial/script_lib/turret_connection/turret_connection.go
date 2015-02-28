package turret_connection

import (
	"github.com/corenomial/script_lib/automationd"
	"github.com/corenomial/script_lib/ssh_connection_to_device"
	"github.com/corenomial/script_lib/tcp_connection_to_device"
	"log"
)

type TurretConnection struct {
	ip    string
	name  string
	creds ssh_connection_to_device.SSHCredentialCreator
	SSH   *ssh_connection_to_device.SSHConnection
	TCP   *tcp_connection_to_device.TCPConnection
}

func ConnectionFactory(ip string, name string, creds ssh_connection_to_device.SSHCredentialCreator) (pconn *TurretConnection) {
	return &TurretConnection{ip, name, creds, nil, nil}
}

func PasswordBasedCredentialsFactory(user_arg, password_arg string) (pcred ssh_connection_to_device.SSHCredentialCreator) {
	return ssh_connection_to_device.PasswordBasedCredentialsFactory(user_arg, password_arg)
}

/***currently broken
func KeyBasedCredentialsFactory(user_arg, key_path string) (pcred ssh_connection_to_device.SSHCredentialCreator) {
	return ssh_connection_to_device.KeyBasedCredentialsFactory(user_arg, key_path)
}***/

func (pconn *TurretConnection) Start() (err error) {
	pconn.SSH = ssh_connection_to_device.ConnectionFactory(pconn.ip, 22, pconn.name)
	err = pconn.SSH.Start(pconn.creds)
	if err != nil {
		log.Fatal("couldn't start SSH connection to ", pconn.name)
	}
	err = pconn.SSH.SendCommands(automationd.CommandsToStartAutomationD, false)
	if err != nil {
		log.Fatal("couldn't start automationd on ", pconn.name)
	}

	pconn.TCP = tcp_connection_to_device.ConnectionFactory(pconn.ip, 8781, pconn.name)
	err = pconn.TCP.Start()
	if err != nil {
		log.Fatal("couldn't start TCP connection to ", pconn.name)
	}
	return
}

func (pconn *TurretConnection) Stop() (err error) {
	if pconn.SSH != nil {
		pconn.SSH.Stop()
	}
	if pconn.TCP != nil {
		pconn.TCP.Stop()
	}
	return
}
