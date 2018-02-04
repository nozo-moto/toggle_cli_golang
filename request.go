package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doError(errorbody string) error {
	return errors.New(errorbody)
}

type StartTimeEntryData struct {
	Created_with string   `json:"created_with"`
	Description  string   `json:"description"`
	Pid          int      `json:"pid"`
	Tags         []string `json:"tags"`
}

type StartTimeEntry struct {
	Time_entry StartTimeEntryData `json:"time_entry"`
}

func MakeStartTimeEntryJson(ste_json StartTimeEntryData) ([]byte, error) {
	st := StartTimeEntry{
		Time_entry: ste_json,
	}
	jsonBytes, err := json.Marshal(st)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return nil, err
	}
	return jsonBytes, nil
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
