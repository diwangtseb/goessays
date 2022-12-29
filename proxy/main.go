package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	ps := &ProxyServer{
		host: "127.0.0.1",
		port: ":2345",
	}
	go ps.Start()
	fmt.Println(http.ListenAndServe(":1234", ps))
}

type ProxyServer struct {
	host string
	port string
}

var _ http.Handler = (*ProxyServer)(nil)

func (ps *ProxyServer) Start() {
	http.HandleFunc("/sh", ps.HandleSayHello)
	http.HandleFunc("/sg", ps.HandleSayGoodBye)
	http.ListenAndServe(ps.port, nil)
}

func (ps *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	proxyUrl, err := url.Parse("http://" + ps.host + ps.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.ServeHTTP(w, r)
}

func (ps *ProxyServer) HandleSayHello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func (ps *ProxyServer) HandleSayGoodBye(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Good, bye!\n")
}
