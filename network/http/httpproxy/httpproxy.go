package httpproxy

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpProxy struct {}


func (*HttpProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		resp *http.Response
		data []byte
		err  error
	)
	r.RequestURI = ""
	r.ParseForm()
	resp, err = http.DefaultClient.Do(r)
	if err != nil {
		http.NotFound(w, r)
		log.Println("1、NotFound")
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		http.NotFound(w, r)
		log.Println("2、NotFound")
		return
	}

	for i, j := range resp.Header {
		for _, k := range j {
			w.Header().Add(i, k)
			log.Println("Header:", i, "=", k)
		}
	}

	for _, c := range resp.Cookies() {
		w.Header().Add("Set-Cookie", c.Raw)
		log.Println("Set-Cookie", c.Raw)
	}
	_, ok := resp.Header["Content-Length"]
	if !ok && resp.ContentLength > 0 {
		w.Header().Add("Content-Length", fmt.Sprint(resp.ContentLength))
		log.Println("1、Content-Length", resp.ContentLength)
	} else {
		log.Println("2、Content-Length", resp.Header["Content-Length"])
	}

	log.Printf("resp.StatusCode:%d  len:%d\n", resp.StatusCode, len(data))
	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}