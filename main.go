package main

import (
	"fmt"
	"github.com/benschw/dns-clb-go/clb"
	"net/http"
	"strings"
	"io/ioutil"
)

func getAddress(svcName string) (string, error) {
	c := clb.NewClb("192.168.101.213", "53", clb.Random)

	srvRecord := svcName + ".service.consul"
	address, err := c.GetAddress(srvRecord)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}



func login() {
	rs, _ := getAddress("station_lugh")
	url :=  fmt.Sprintf("http://%s/api/v1/public/staff/sso/login", rs)
	fmt.Println(url)
	payload := strings.NewReader("{\n  \"username\":\"tamnt70\",\n  \"password\":\"admin1235\"\n\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "9a73b716-36a5-8672-42f4-40ccfad283ac")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}


func main() {
	//rs, err := getAddress("station_lugh")
	login()
	//fmt.Println(rs, err)
}


