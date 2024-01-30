package pal

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gorcon/rcon"
)

type PalClient struct {
	conn       *rcon.Conn
	ipAndPort  string
	password   string
	retryDelay int
	retryCount int
}

func NewPalClient(ip string, port int, password string, retryDelay int, retryCount int) *PalClient {
	ipAndPort := ip + ":" + strconv.Itoa(port)
	return &PalClient{
		ipAndPort:  ipAndPort,
		password:   password,
		retryDelay: retryDelay,
		retryCount: retryCount,
	}
}

func (client *PalClient) Connect() error {
	fmt.Println("ipAndPort=" + client.ipAndPort + ", pwd=" + client.password)
	conn, err := rcon.Dial(client.ipAndPort, client.password)
	if err != nil {
		return fmt.Errorf("Failed to connect to RCON server: %v", err)
	}
	client.conn = conn
	fmt.Println("RCON connected")
	return nil
}

func (client *PalClient) Close() {
	if client.conn == nil {
		return
	}
	client.conn.Close()
	client.conn = nil
}

func (client *PalClient) execute(command string) (string, error) {
	resp, err := client.conn.Execute(command)
	if err != nil {
		return "", fmt.Errorf("Failed to send RCON command: %v", err)
	}
	fmt.Println("RCON response:", resp)
	return resp, nil
}

func (client *PalClient) sendCommand(command string, retryCount int) (string, error) {
	if client.conn == nil {
		if retryCount > 0 {
			if client.retryDelay > 0 {
				time.Sleep(time.Duration(client.retryDelay) * time.Second)
			}
			fmt.Println("no conn, retryCount=" + strconv.Itoa(retryCount))
			client.Connect()
			return client.sendCommand(command, retryCount-1)
		} else {
			return "", fmt.Errorf("RCON connection is not established")
		}
	}
	resp, err := client.execute(command)
	if err != nil {
		if retryCount > 0 {
			fmt.Println("send command faild, retryCount="+strconv.Itoa(retryCount), err)
			time.Sleep(time.Duration(client.retryDelay) * time.Second)
			client.Close()
			client.Connect()
			return client.sendCommand(command, retryCount-1)
		} else {
			return "", err
		}
	}
	return resp, nil
}

func (client *PalClient) ShowPalyers() (string, error) {
	return client.sendCommand("ShowPlayers", client.retryCount)
}

func (client *PalClient) Info() (string, error) {
	return client.sendCommand("Info", client.retryCount)
}

func (client *PalClient) Broadcast(message string) (string, error) {
	return client.sendCommand(`Broadcast `+message, client.retryCount)
}

func (client *PalClient) Reboot(message string) (string, error) {
	return client.sendCommand(`Shutdown `+message+` server_will_reboot`, client.retryCount)
}
