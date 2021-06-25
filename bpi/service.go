package bpi

import (
	"encoding/json"
	"fmt"
	"proxyblockpi/network/http/httpclient"
)

type Provider struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Logo  string `json:"logo"`
	Website  string `json:"website"`
}

type ServiceEndPoint struct {
	Endpoint       string `json:"endpoint"`
	Weight         int `json:"weight"`
}
type ServiceEndPoints []ServiceEndPoint

type Service struct {
	SID          string `json:"sid"`
	Name         string `json:"name"`
	Provider     Provider `json:"provider"`
	AccessKey    string `json:"access_key"`
	Endpoints    ServiceEndPoints `json:"endpoints"`
}
type Services []Service

type Result struct {
	Code  	string `json:"code"`
	Data  	Services `json:"data"`
	Message string `json:"message"`
}


var GServices = map[string]Service{}

func Book() {

}

func LoadServices() {
	//http://192.168.101.215:8000/api/subscribe

	r := httpclient.Post("http://192.168.101.215:8000/api/subscribe", "sids=c75d5e926d144a54e77afd94b057cb4c512535bf,45bc68784ddedb4177387ffec9dbcad54a287977")
	result := Result{}
	if err := json.Unmarshal([]byte(r), &result); err != nil {
		fmt.Println("Error LoadServices: ", err)
		return
	}

	if result.Code != "SUCCESS" {
		fmt.Println("Error LoadServices: ", result.Code)
		return
	}

	for _, _service := range result.Data {
		GServices[_service.SID]=_service
		fmt.Println("Loaded service: SID ", _service.SID, " - ", _service.Name)
	}

	println(r)
}