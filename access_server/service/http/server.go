package http_service

import (
    "crypto/tls"
    "log"
    "net/http"

    "online_consultant/server/access_server/config"
)

// Service handler register
var handlerFunc map[string]http.HandlerFunc

func RegisterServiceHandler(name string, function http.HandlerFunc) {
    if (handlerFunc == nil) {
        handlerFunc = make(map[string]http.HandlerFunc)
    }
    handlerFunc[name] = function
}

// Http server
type HttpServer struct {
    server *http.Server
    handler *http.ServeMux
}

func NewServer() (*HttpServer) {
    server := new(HttpServer)
    // Init service
    InitRelayService()
    // build http handler
    server.handler = http.NewServeMux()
    for name, function := range handlerFunc {
        server.handler.HandleFunc(name, function)
    }
    // new http server
    server.server = &http.Server{
        Addr: config.GetConfigManager().Addr,
        Handler: server.handler,
        // Disable HTTP/2.
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
    }
    return server
}

func (this *HttpServer) Start() {
    configManager := config.GetConfigManager() 
    if configManager.PemPath != "" && configManager.KeyPath != "" {
        this.StartHttpsServer(configManager.PemPath, configManager.KeyPath)
    } else {
        this.StartHttpServer()
    }
}

func (this *HttpServer) StartHttpServer() {
    log.Fatal(this.server.ListenAndServe())
}

func (this *HttpServer) StartHttpsServer(pemPath string, keyPath string) {
	log.Fatal(this.server.ListenAndServeTLS(pemPath, keyPath))
}
