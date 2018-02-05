package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

type CurrentRespData struct {
	Id          int    `json:"id"`
	Wid         int    `json:"wid"`
	Uid         int    `json:"uid"`
	Pid         int    `json:"pid"`
	Billiable   bool   `json:"billable"`
	Start       string `json:"start"`
	Guid        string `json:"guid"`
	Duronly     bool   `json:"duronly"`
	Duration    int    `json:"duration"`
	Description string `json:"description"`
	At          string `json:"at"`
}
type CurrentResp struct {
	Data CurrentRespData `json:"data"`
}

func current_toggle() (*CurrentResp, error) {
	url := "https://www.toggl.com/api/v8/time_entries/current"
	request_data := Request{
		method:   "GET",
		url:      url,
		jsondata: nil,
	}
	result_byte, err := request(
		request_data,
	)
	if err != nil {
		fmt.Println("error ", err)
		return nil, err
	}

	result := new(CurrentResp)
	if err := json.Unmarshal(result_byte, result); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}
	return result, err
}

type StopRespData struct {
	Id          int      `json:"id"`
	Pid         int      `json:"pid"`
	Wid         int      `json:"wid"`
	Billiable   bool     `json:"billable"`
	Start       string   `json:"start"`
	Duration    int      `json:"duration"`
	Description string   `json:"description"`
	tags        []string `json:"tags"`
}
type StopResp struct {
	Data StopRespData `json:"data"`
}

func stop_toggle(time_entry_id int) (*StopResp, error) {
	url := fmt.Sprint("https://www.toggl.com/api/v8/time_entries/", strconv.Itoa(time_entry_id), "/stop")
	request_data := Request{
		method:   "PUT",
		url:      url,
		jsondata: nil,
	}
	result_byte, err := request(
		request_data,
	)
	if err != nil {
		return nil, err
	}
	result := new(StopResp)
	if err := json.Unmarshal(result_byte, result); err != nil {
		return nil, err
	}
	return result, err
}

type StartRespData struct {
	Id          int      `json:"id"`
	Pid         int      `json:"pid"`
	Wid         int      `json:"wid"`
	Billiable   bool     `json:"billable"`
	Start       string   `json:"start"`
	Duration    int      `json:"duration"`
	Description string   `json:"description"`
	tags        []string `json:"tags"`
}
type StartResp struct {
	Data StartRespData `json:"data"`
}

func start_toggle() (*StopResp, error) {
	url := "https://www.toggl.com/api/v8/time_entries/start"
	ste_data := StartTimeEntryData{
		Created_with: "golang",
		Description:  "Meeting with possible clients",
		Tags:         []string{"billed"},
	}
	data, err := MakeStartTimeEntryJson(ste_data)
	if err != nil {
		return nil, err
	}
	request_data := Request{
		method:   "POST",
		url:      url,
		jsondata: data,
	}
	result_byte, err := request(
		request_data,
	)
	if err != nil {
		return nil, err
	}

	result := new(StopResp)
	if err := json.Unmarshal(result_byte, result); err != nil {
		return nil, err
	}
	return result, err
}

func stop() error {
	current_result, err := current_toggle()
	if err != nil {
		return err
	}
	_, err = stop_toggle(current_result.Data.Id)
	if err != nil {
		return err
	}
	return nil
}

func start() error {
	current_result, err := current_toggle()
	if err != nil {
		return err
	}
	if len(current_result.Data.Start) > 0 {
		return doError("Toggle is Running")
	}
	_, err = start_toggle()
	if err != nil {
		return err
	}
	return nil
}
