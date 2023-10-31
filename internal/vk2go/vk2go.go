package vk2go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type LastPublicationsResponse struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			Id   int 	`json:"id"`
			Text string `json:"text"`
		} 
	} `json:"response"`
}

func GetDataFromVk(token string, reqVersion float32, domain string, count int) (*LastPublicationsResponse, error) {
	url := fmt.Sprintf(
		"https://api.vk.com/method/wall.get?access_token=%s&v=%.3f&domain=%s&count=%d",
		token, reqVersion, domain, count,
	)

	var client http.Client
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()


	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(res.Status)
	}

 	bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        logrus.Fatal(err)
    }

	var ret LastPublicationsResponse
	if err := json.Unmarshal(bodyBytes, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
