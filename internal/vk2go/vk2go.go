package vk2go

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)


func GetDataFromVk(token, domain string, reqVersion float32) (string, error) {
	logrus.Infof("Trying to get posts from VK for https://vk.com/%s", domain)

	url := fmt.Sprintf(
		"https://api.vk.com/method/wall.get?access_token=%s&v=%.3f&domain=%s",
		token, reqVersion, domain,
	)

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
	    logrus.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
 		bodyBytes, err := io.ReadAll(resp.Body)
    	if err != nil {
    	    logrus.Fatal(err)
    	}
    	bodyString := string(bodyBytes)
		return bodyString, nil
	}
	return "", nil
}
