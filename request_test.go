package main

import (
	"testing"
)

func TestMakeStartTimeEntryJson(t *testing.T) {
	ste_json, err := MakeStartTimeEntryJson(
		"golang",
		"Test From Golang",
		123,
		[]string{"billed"},
	)
	if err != nil {
		t.Fatal("failed test", err)
	}
	result_json := `{"time_entry":{"create_with":"golang","description":"Test From Golang","pid":123,"tags":["billed"]}}`
	if ste_json != result_json {
		t.Fatal("failed test", ste_json)
	}
}
