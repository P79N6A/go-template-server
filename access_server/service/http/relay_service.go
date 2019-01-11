// Copyright 2018 doctorwechat
//
// Author: juzhongguoji <juzhongguoji@hotmail.com>
// Date:   2018/12/18

package http_service

import (
	"io"
	//"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"strings"

    "online_consultant/server/access_server/config"

    glog "github.com/golang/glog"
    //"github.com/golang/protobuf/jsonpb"
)

//var jsonMarshaler = jsonpb.Marshaler{EmitDefaults: true}

var relayService RelayService
const (
	urlPathSegmentNum = 5 // /APP/Server/Service/Function -> ['', 'APP', 'Server', 'Service', 'Function']
)

type RelayService struct {
	balancer map[string]roundRobinPicker
}

func InitRelayService() {
	relayService.initBalancer()

	RegisterServiceHandler("/", 
		func(response http.ResponseWriter, request *http.Request) {
			relayService.RelayRequest(response, request)
		},
	)
}

func (this *RelayService) initBalancer() {
	this.balancer = make(map[string]roundRobinPicker)
	routerTable := config.GetRouterTable()
	for name, addrs := range routerTable.Addrs {
		this.balancer[name] = roundRobinPicker{
			list: addrs,
			next: 0,
		}
	}	
}

func (this *RelayService) RelayRequest(response http.ResponseWriter, request *http.Request) {
	glog.V(2).Info("Get request: ",request.URL)
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	glog.V(2).Info("Parse request")
	requestHeader, requestData, err := this.parseRequest(requestBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	glog.V(2).Info("Get relay address")
	relayAddress := this.getRelayAddress(request.URL) 
	glog.Infoln("Relay address: ", relayAddress)
	client := &http.Client{}
	relayRequest, err := http.NewRequest("POST", 
			relayAddress,
			request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	glog.V(2).Info("Relay request")
	relayResponse, err := client.Do(relayRequest)
	if err != nil {
		http.Error(response, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer relayResponse.Body.Close()
	//this.copyHeader(response.Header(), relayResponse.Header)
	response.WriteHeader(relayResponse.StatusCode)
	// TODO: 转换为外部请求头
	io.Copy(response, relayResponse.Body)
	glog.V(2).Info("Send response")
}

func (this *RelayService) parseRequest(requestBody string) (string, string, error) {
	routerName := this.getRouterName(reqUrl.Path)
	glog.V(2).Info(routerName)
	picker, found := this.balancer[routerName]
	if (!found) {
		return ""
	}
	reqUrl.Host = picker.Pick()
	
	if (reqUrl.Scheme == "") {
		reqUrl.Scheme = "http"
	}

	return reqUrl.String()
}

func (this *RelayService) getRelayAddress(reqUrl *url.URL) string {
	routerName := this.getRouterName(reqUrl.Path)
	glog.V(2).Info(routerName)
	picker, found := this.balancer[routerName]
	if (!found) {
		return ""
	}
	reqUrl.Host = picker.Pick()
	
	if (reqUrl.Scheme == "") {
		reqUrl.Scheme = "http"
	}

	return reqUrl.String()
}

func (this *RelayService) copyHeader(dst, src http.Header) {
    for k, vv := range src {
        for _, v := range vv {
            dst.Add(k, v)
        }
    }
}

func (this *RelayService) getRouterName(urlPath string) string {
	segments := strings.Split(urlPath, "/")
	if len(segments) != urlPathSegmentNum {
		glog.V(2).Infoln("Invalid url, path segment num ",len(segments), 
				"(expect ",urlPathSegmentNum,")")
		return ""
	}

	return segments[1] + "." + segments[2] + "." + segments[3]
}

func (this *RelayService) accessHeaderToRequestHeader() {

}

func (this *RelayService) responseHeaderToAccessHeader() {

}

type roundRobinPicker struct {
	list []string
	mutex sync.Mutex
	next int
}

func (picker *roundRobinPicker) Pick() (string) {
	if len(picker.list) <= 0 {
		return ""
	}

	picker.mutex.Lock()
	item := picker.list[picker.next]
	picker.next = (picker.next + 1) % len(picker.list)
	picker.mutex.Unlock()
	return item
}