package main

import (
	"testing"
	"time"
)

func TestParsingStringOverToDate(t *testing.T) {
	date := getLastEditTime()
	date_string := getLastEditTime().String()

	date_from_string, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", date_string)
	if err != nil {
		println("Error parsing date..")
		println(err.Error())
	}
	if date.Equal(date_from_string) {
		println("Test runned nicely.")
	}else {
		t.Failed()
	}

}