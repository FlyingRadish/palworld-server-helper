package rcn

import (
	"fmt"
	"strconv"

	"github.com/FlyingRadish/rcong"
)

type RCNClient struct {
	conn       *rcong.RCONConnection
	ip         string
	port       int
	ipAndPort  string
	password   string
	retryDelay int
	retryCount int
}

var instance RCNClient

func Create(ip string, port int, password string, retryCount int, retryDelay int) {
	ipAndPort := ip + ":" + strconv.Itoa(port)
	conn := rcong.NewRCONConnection(ip, port, password, retryCount, retryDelay)
	conn.Connect()
	instance = RCNClient{
		conn:       conn,
		ip:         ip,
		port:       port,
		ipAndPort:  ipAndPort,
		password:   password,
		retryDelay: retryDelay,
		retryCount: retryCount,
	}
}

func GetRCNClient() *RCNClient {
	return &instance
}

func (client *RCNClient) Close() {
	if client.conn == nil {
		return
	}
	client.conn.Close()
}

func (client *RCNClient) sendCommand(command string) (string, error) {
	resp, err := client.conn.ExecCommand(command)
	if err != nil {
		return "", fmt.Errorf("Failed to send RCON command: %v", err)
	}
	// fmt.Println("RCON response:", resp)
	return resp, nil
}

func (client *RCNClient) ShowPalyers() (string, error) {
	return client.sendCommand("ShowPlayers")
}

func (client *RCNClient) Info() (string, error) {
	return client.sendCommand("Info")
}

func (client *RCNClient) Broadcast(message string) (string, error) {
	return client.sendCommand(`Broadcast ` + message)
}

func (client *RCNClient) Reboot(message string) (string, error) {
	return client.sendCommand(`Shutdown ` + message + ` server_will_reboot`)
}

func (client *RCNClient) ExecCommand(command string) (string, error) {
	return client.sendCommand(command)
}
