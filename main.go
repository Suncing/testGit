package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getTokenBaidu() {

	var host = "https://aip.baidubce.com/oauth/2.0/token"
	var param = map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     "y3XGnsTwiMvwve5FExqcMc3T",
		"client_secret": "9tj3ZVafjTkYG2p2jQk5wBuK9CSkt6oE",
	}

	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	for k, v := range param {
		query.Set(k, v)
	}
	uri.RawQuery = query.Encode()

	response, err := http.Get(uri.String())
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}

func CheckImage() {
	var host = "https://aip.baidubce.com/rest/2.0/solution/v1/img_censor/v2/user_defined"
	var accessToken = "24.16f06da3aad3e7bd2209adc57cf5b793.2592000.1665736497.282335-27465655"
	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()
	filebytes, err := ioutil.ReadFile("C:\\Users\\nova005344\\Desktop\\123.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendData := sendBody.Form.Encode()
	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
