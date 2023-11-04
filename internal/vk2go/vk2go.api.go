package vk2go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	FETCH_COUNT = 5
)

var ErrNoNewPublications = errors.New("no new publications since last update")

func NewPublications(storage PublicationsStorage, domain string, token string, reqVersion float32) ([]Publication, error) {
	data, err := fetchDataFromVk(domain, FETCH_COUNT, token, reqVersion)
	if err != nil {
		return nil, err
	}
	var publications []Publication
	for _, item := range data.Response.Items {
		published, err := storage.IsPublished(item.Id, domain)
		if err != nil {
			logrus.Errorf("error checking already published: %s", err.Error())
		}
		if !published {
			publications = append(publications, Publication{
				Id: item.Id,
				Text: item.Text,
			})
			err = storage.MarkAsPublished(item.Id, domain)
			if err != nil {
				logrus.Errorf("error marking as published: %s", err.Error())
			}
		}
	}
	if publications == nil {
		return nil, ErrNoNewPublications
	}
	return publications, nil
}

type lastPublicationsData struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			Id   int 	`json:"id"`
			Text string `json:"text"`
		} 
	} `json:"response"`
}

func fetchDataFromVk(domain string, count int, token string, reqVersion float32) (*lastPublicationsData, error) {
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

	var ret lastPublicationsData
	if err := json.Unmarshal(bodyBytes, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
