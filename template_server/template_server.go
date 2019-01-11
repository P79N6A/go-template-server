package main 

import (
    "flag"
    "net/http"
    "strings"

    confmgr "online_consultant/server/template_server/config"
    httpservice "online_consultant/server/template_server/service/http_service"
    rpcservice "online_consultant/server/template_server/service/rpc_service"
    svrresolver "online_consultant/server/template_server/service/name_resolver"

    glog "github.com/golang/glog"
    "golang.org/x/net/http2"
    "golang.org/x/net/http2/h2c"
)

const (
    defaultConfigFilePath string = "../config/template_server.conf"
)

func init() {
    configManager := confmgr.GetInstance()
    flag.StringVar(&configManager.ConfigFilePath, "config", defaultConfigFilePath, "Config file path")
    flag.Parse()
}

func main() {
    // Load config
    configManager := confmgr.GetInstance()
    if configManager.Init() == false {
        glog.Fatal("Load server config failed!")
    }
    glog.Info("Load server config success")
    // Load resolver config
    if svrresolver.GetResolverConfig().Init(configManager.ResolverConfigFilePath) == false {
        glog.Fatal("Load resolver config failed!")
    }
    glog.Info("Load resolver config success")
    
    // Create servers
    httpHandler := httpservice.CreateHandler()
    rpcServer := rpcservice.CreateServer()
    // Start server
    glog.Info("Start server...")
    http.ListenAndServe(configManager.Addr,
        h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
                rpcServer.ServeHTTP(w, r)
            } else {
                httpHandler.ServeHTTP(w, r)
            }
        }), &http2.Server{}),
    )
}
