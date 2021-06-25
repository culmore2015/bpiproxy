package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"proxyblockpi/bpi"
	"strings"
)


type handler struct {
}

func ParseSIDAndKey(r *http.Request) (string,string) {
	val := strings.Split(r.URL.Path[1:],"/")
	if len(val) < 2 || len(val[0]) != 40 || len(val[1]) != 40{
		panic("Invalid API Path: " + r.URL.Path)
	}
	return val[0],val[1]
}


func CopyResponse(rw http.ResponseWriter, resp *http.Response) error {
	for k, vv := range resp.Header {
		for _, v := range vv {
			rw.Header().Add(k, v)
		}
	}
	rw.WriteHeader(resp.StatusCode)
	defer resp.Body.Close()
	_, err := io.Copy(rw, resp.Body)
	return err
}

func (*handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		if err:=recover(); err!=nil {
			fmt.Println("Panic: ", err)
		}
	}()
	fmt.Println("=========New request=========")
	SID,KEY := ParseSIDAndKey(r)
	service,exist := bpi.GServices[SID]
	if !exist {
		panic("Service NOT exist: " + SID)
	}

	fmt.Println("Service: ", service.Name + "(" + service.SID + ")")
	fmt.Println("Access Key: ", KEY)

	endpoint := service.Endpoints[0].Endpoint
	fmt.Println("Endpoint: ", endpoint)

	fmt.Println("Request Method: ", r.Method)

	transport := http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *r
	if clientIP,_,err := net.SplitHostPort(r.RemoteAddr); err==nil{
		if prior, ok := outReq.Header["For"]; ok {
			clientIP = strings.Join(prior,",") + ", " + clientIP
		}
		outReq.Header.Set("For", clientIP)
	}

	URI := strings.Replace(r.RequestURI, "/"+SID+"/"+KEY,"", -1)
	outURL,_ := url.Parse(endpoint + URI)
	*outReq.URL = *outURL
	outReq.Host = outReq.URL.Host
	outReq.RequestURI =  outReq.URL.Path

	resp, err :=transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		fmt.Println("Request Error: ", err)
		return
	}

	CopyResponse(rw, resp)


	//u, _ := url.Parse(endpoint)
	//ws.NewProxy(u).ServeHTTP(rw, r)
}



