// Copyright 2018 doctorwechat
//
// Author: juzhongguoji <juzhongguoji@hotmail.com>
// Date:   2018/12/22

package config

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "regexp"
)

const serverConfigPath = "../config/access_server.conf"
var configManager ConfigManager

func GetInstance() *ConfigManager {
    return &configManager
}

type ConfigManagerInterface interface {
    Init() bool
}

type ConfigManager struct {
    ConfigFilePath string
    Addr string // ip:port
    ResolverConfigFilePath string
}

func (this *ConfigManager) Init() bool {
    content, e := ioutil.ReadFile(this.ConfigFilePath)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        return false
    }
    
    content, e = stripJsonComments(content)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        return false
    }

    fmt.Printf("Read config file: %s\n", this.ConfigFilePath)
    json.Unmarshal(content, &this)
    fmt.Printf("Results: %v\n", this)

    return true
}

func stripJsonComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0) // Windows
	lines := bytes.Split(data, []byte("\n"))                //split to muli lines
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}

	return bytes.Join(filtered, []byte("\n")), nil
}