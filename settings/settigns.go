package settings

import (
	"encoding/json"
	"fmt"
	"os"
)

type HelperConfig struct {
	IP                  string  `json:"ip"`
	Port                int     `json:"port"`
	Password            string  `json:"password"`
	RetryCount          int     `json:"retryCount"`
	RetryDelay          int     `json:"retryDelay"`
	RebootScriptPath    string  `json:"rebootScriptPath"`
	RebootSeconds       int     `json:"rebootSeconds"`
	OOMThreshold        float64 `json:"oomThreshold"`
	OOMCheckInterval    int     `json:"oomCheckIntervalSeconds"`
	PlayerCheckInterval int     `json:"playerCheckInterval"`
	ApiHost             string  `json:"apiHost"`
	ApiPort             int     `json:"apiPort"`
}

var configPath = "./helper_config.json"

func SetConfigPath(path string) {
	configPath = path
}

func LoadConfig() (HelperConfig, error) {
	file, err := os.Open(configPath)
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
	fmt.Println(config)
	return config, nil
}
