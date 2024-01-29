package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
	"encoding/json"
	"os"
	"strconv"

	"pal-server-helper/pal"
	"github.com/shirou/gopsutil/mem"
)

type HelperConfig struct {
	ServerIPAndPort 	 string `json:"serverIPAndPort"`
	ServerPassword 	 	 string `json:"serverPassword"`
	RebootScriptPath     string `json:"rebootScriptPath"`
	OOMThreshold    	float64 `json:"oomThreshold"`
	Interval         		int `json:"checkIntervalSeconds"`
	RebootSeconds       	int `json:"rebootSeconds"`
}

func main() {
	fmt.Println("pal server helper started")
	fmt.Println("loading config...")
	config, err := loadConfig("helper_config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fmt.Println("Connect to server...")
	client := pal.NewPalClient()
	connectErr := client.Connect(config.ServerIPAndPort, config.ServerPassword)
	if connectErr != nil {
		fmt.Println(connectErr)
		return
	}
	defer client.Close()

	go monitorMemoryUsage(client, config)

	// 主程序持续执行
	for {
		time.Sleep(10 * time.Second) // 可以根据实际需求调整间隔时间
		// 在这里可以添加其他持续执行的逻辑
	}

	// 在程序即将退出时执行client.Close()
	defer client.Close()
}


func loadConfig(filename string) (HelperConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return HelperConfig{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := HelperConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		return HelperConfig{}, err
	}

	return config, nil
}

func monitorMemoryUsage(client *pal.PalClient, config HelperConfig) {
	ticker := time.NewTicker(time.Duration(config.Interval) * time.Second)
	defer ticker.Stop()

	var wg sync.WaitGroup

	for {
		select {
		case <-ticker.C:
			memory, err := mem.VirtualMemory()
			if err != nil {
				fmt.Println("Failed to get memory info:", err)
				continue
			}

			usedPercent := memory.UsedPercent
			fmt.Printf("Current memory usage: %.2f%%\n", usedPercent)
			if usedPercent > config.OOMThreshold {
				wg.Add(2)
				go notifyReboot(&wg, client, config)
				go reboot(&wg, config)
				wg.Wait() // 阻塞，直到重启完成
			}
		}
	}
}

func notifyReboot(wg *sync.WaitGroup, client *pal.PalClient, config HelperConfig) {
	defer wg.Done()
	// 执行通知重启操作
	fmt.Println("Notify to reboot")
	seconds := strconv.Itoa(config.RebootSeconds)
	client.Reboot(seconds)
	countdown := config.RebootSeconds
	for countdown > 0 {
		message := fmt.Sprintf("OOM_server_reboot_in_%ds", countdown)
		client.Broadcast(message)
		time.Sleep(time.Second)
		countdown--
	}
}


func reboot(wg *sync.WaitGroup, config HelperConfig) {
	defer wg.Done()
	fmt.Println("Waiting to restart...")
	time.Sleep(10 * time.Second)
	// 执行重启操作
	fmt.Println("Rebooting...")
	cmd := exec.Command("sh", config.RebootScriptPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to execute reboot script:", err)
	} else {
		fmt.Println("Reboot script executed successfully")
		// 阻塞2分钟，等待服务重启
		fmt.Println("Waiting for service to restart...")
		time.Sleep(2 * time.Minute)
		fmt.Println("Service restarted")
	}
}
