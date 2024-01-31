package main

import (
	"flag"
	"fmt"
	"time"

	"pal-server-helper/api"
	"pal-server-helper/common"
	"pal-server-helper/pal"
	"pal-server-helper/pal/rcn"
	"pal-server-helper/settings"
)

func main() {
	// 定义命令行参数
	configFilePath := flag.String("c", "", "Path to the configuration file")
	flag.Parse()
	var configPath = ""
	if *configFilePath == "" {
		configPath = "helper_config.json"
		fmt.Println("use default config path=" + configPath)
	} else {
		configPath = *configFilePath
	}
	settings.SetConfigPath(configPath)

	fmt.Println("pal server helper started")
	fmt.Println("loading config...")
	config, err := settings.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	api.UpdatePanelApi(config.ApiHost, config.ApiPort, config.PanelPath)
	fmt.Println("Connect to server...")
	rcn.Create(config.IP, config.Port, config.Password, config.RetryCount, config.RetryDelay)
	client := rcn.GetRCNClient()
	fmt.Println("Connected")
	defer client.Close()

	go common.MonitorMemoryUsage(config.OOMThreshold, config.OOMCheckInterval)
	go pal.MonitorPlayers(config.OOMCheckInterval)
	go api.RunApiServer(config.ApiPort, config.PanelPath)

	// 主程序持续执行
	for {
		time.Sleep(10 * time.Second) // 可以根据实际需求调整间隔时间
		// 在这里可以添加其他持续执行的逻辑
	}
}
