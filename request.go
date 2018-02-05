package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doError(errorbody string) error {
	return errors.New(errorbody)
}

type Request struct {
	method   string
	url      string
	jsondata []byte
}

func request(data Request) ([]byte, error) {
	req, err := http.NewRequest(
		data.method,
		data.url,
		bytes.NewBuffer([]byte(data.jsondata)),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(API_TOKEN, "api_token")
	client := new(http.Client)
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else if resp.Status != "200 OK" {
		return nil, doError(fmt.Sprint(resp.Status, string(body)))
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}
