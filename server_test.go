package main

import (
	"testing"
	"time"
)

/**
	To check if our parse actually works..
 */
func TestParsingStringOverToDate(t *testing.T) {
	date := getLastEditTime()
	// get the same date, and turn it over to string.. because we get time in string format from api..
	date_string := getLastEditTime().String()

	// parsing date from string to time format.
	date_from_string, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", date_string)
	if err != nil {
		println("Error parsing date..")
		t.Failed()
		println(err.Error())
	}
	// checking if the date is the same.. it should be..
	if date.Equal(date_from_string) {
		println("Test runned nicely.")
	}else {
		t.Failed()	// test failed
	}
}