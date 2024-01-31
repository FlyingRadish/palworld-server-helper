package pal

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"pal-server-helper/pal/rcn"
	"pal-server-helper/settings"
)

func Reboot(needNotify bool) error {
	var wg sync.WaitGroup
	config, err := settings.LoadConfig()
	if err != nil {
		return errors.New("reboot failed because config not loaded")
	}
	client := rcn.GetRCNClient()
	if needNotify {
		wg.Add(2)
		notifyReboot(&wg, client, config.RebootSeconds)
		reboot(&wg, client, config.RebootScriptPath)
		wg.Wait() // 阻塞，直到重启完成
	} else {
		wg.Add(1)
		reboot(&wg, client, config.RebootScriptPath)
		wg.Wait() // 阻塞，直到重启完成
	}
	return nil

}

func notifyReboot(wg *sync.WaitGroup, client *rcn.RCNClient, rebootSeconds int) {
	defer wg.Done()
	// 执行通知重启操作
	fmt.Println("Notify to reboot")
	seconds := strconv.Itoa(rebootSeconds)
	client.Reboot(seconds)
	countdown := rebootSeconds
	for countdown > 0 {
		message := fmt.Sprintf("OOM_server_reboot_in_%ds", countdown)
		client.Broadcast(message)
		time.Sleep(time.Second)
		countdown--
	}
}

func reboot(wg *sync.WaitGroup, client *rcn.RCNClient, rebootScriptPath string) {
	defer wg.Done()
	fmt.Println("closing RCON client")
	client.Close()
	fmt.Println("Waiting to restart...")
	time.Sleep(10 * time.Second)
	// 执行重启操作
	fmt.Println("Rebooting...")
	cmd := exec.Command("sh", rebootScriptPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to execute reboot script:", err)
	} else {
		fmt.Println("Reboot script executed successfully")
		client.Close()
		// 阻塞2分钟，等待服务重启
		fmt.Println("Waiting for service to restart...")
		time.Sleep(2 * time.Minute)
		fmt.Println("Service restarted")
	}
}

func Backup() error {
	return nil

}
