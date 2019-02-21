package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	rr "github.com/lflxp/api-gw/Utils"
)

var RR = rr.NewWeightedRR(rr.RR_NGINX)

type handle struct {
	addrs []string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addr := RR.Next().(string)
	remote, err := url.Parse("http://" + addr)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//被代理的服务器host和port
	h := &handle{}
	h.addrs = []string{"192.168.50.216:8080"}

	w := 1
	for _, e := range h.addrs {
		RR.Add(e, w)
		w++
	}
	err := http.ListenAndServe(":28080", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
