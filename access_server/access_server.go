package main 

import (
    "flag"

    glog "github.com/golang/glog"

    "online_consultant/server/access_server/config"
    httpservice "online_consultant/server/access_server/service/http"
)

const (
    defaultConfigFilePath string = "../config/access_server.conf"
)

func init() {
    configManager := config.GetConfigManager()
    flag.StringVar(&configManager.ConfigFilePath, "config", defaultConfigFilePath, "Config file path")
    flag.Parse()

    // Load config
    if !configManager.Init() {
        glog.Fatal("Load server config failed!")
    }
    glog.Info("Load server config success")
    // Load router table
    routerTable := config.GetRouterTable()
    if !routerTable.Init(configManager.RouterTableFilePath) {
        glog.Fatal("Load router table failed!")
    }
    glog.Info("Load router table success")
}

func main() {
    // Create servers
    httpServer := httpservice.NewServer()
    // Start server
    glog.Info("Start server...")
    httpServer.Start()
}
