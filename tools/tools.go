package tools

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
	Sms *SmsConfig `json:"smg"`
}


type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateName string `json:"template_name"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RegineId     string `json:"regine_id"`
}


func ParseJson(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	var config Config
	decoder.Decode(&config)

	return &config
}

func JsonParseTOStruct(io io.ReadCloser, obj interface{}) error {
	err := json.NewDecoder(io).Decode(obj)

	return err
}


