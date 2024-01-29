package pal

import (
	"fmt"
	"github.com/gorcon/rcon"
)


type PalClient struct {
    conn *rcon.Conn
}

func NewPalClient() *PalClient {
    return &PalClient{}
}

func (client *PalClient) Connect(ip string, password string) error {
    conn, err := rcon.Dial(ip, password)
    if err != nil {
        return fmt.Errorf("Failed to connect to RCON server: %v", err)
    }
    client.conn = conn
    return nil
}

func (client *PalClient) Close() {
    if client.conn == nil {
        return
    }
    client.conn.Close()
    client.conn = nil
}

func (client *PalClient) Broadcast(message string) error {
    if client.conn == nil {
        return fmt.Errorf("RCON connection is not established")
    }
    resp, err := client.conn.Execute(`Broadcast ` + message)
    if err != nil {
        return fmt.Errorf("Failed to send RCON command: %v", err)
    }
    fmt.Println("RCON response:", resp)
    return nil
}

func (client *PalClient) Reboot(message string) error {
    if client.conn == nil {
        return fmt.Errorf("RCON connection is not established")
    }
    resp, err := client.conn.Execute(`Shutdown `+ message +` server_will_reboot`)
    if err != nil {
        return fmt.Errorf("Failed to send RCON command: %v", err)
    }
    fmt.Println("RCON response:", resp)
    return nil
}