package httpclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)


func Get(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}


func GetResponse(url string) *http.Response {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}



func HttpPost(url string, data string, contentType string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func Post(url string, data string) string {
	return HttpPost(url, data,"application/x-www-form-urlencoded")
}

func PostJson(url string, data interface{}) string {
	jsonStr, _ := json.Marshal(data)
	return HttpPost(url, string(jsonStr),"application/x-www-form-urlencoded")
}