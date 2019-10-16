package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IpInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func main() {

	response, err := http.Get("https://ipinfo.io")
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
	}

	var data IpInfo
	if err := json.Unmarshal(responseData, &data); err != nil {
		log.Print(err)
	}
	fmt.Println("The object is")
	fmt.Printf("%+v\n", data)

	fmt.Println("Your ip is:", data.IP)
	fmt.Println("Your country is:", data.Country)

}
