// Copyright 2018 doctorwechat
//
// Author: juzhongguoji <juzhongguoji@hotmail.com>
// Date:   2018/12/22

package config 

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "regexp"

    glog "github.com/golang/glog"
)

var routerTable RouterTable

func GetRouterTable() *RouterTable{
    return &routerTable
}

// struct & interface define
type RouterTableInterface interface {
    Init() bool
    FindAddressByName(target string) ([]string, bool)
}

type RouterTable struct {
    ConfigFilePath string
    Addrs map[string][]string
}

func (this *RouterTable) Init(configFilePath string) bool {
    this.ConfigFilePath = configFilePath

    content, e := ioutil.ReadFile(configFilePath)
    if e != nil {
        glog.Error("File error: ", e)
        return false
    }
    
    content, e = this.stripJsonComments(content)
    if e != nil {
        glog.Error("File error: ", e)
        return false
    }

    json.Unmarshal(content, &this)
    glog.Info(this)

    return true
}

func (this *RouterTable) FindAddressByName(target string) ([]string, bool) {
    addrs, found := this.Addrs[target]
    return addrs, found
}

func (this *RouterTable) stripJsonComments(data []byte) ([]byte, error) {
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