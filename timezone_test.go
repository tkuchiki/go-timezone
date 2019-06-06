package timezone

import (
	"testing"
)

var timezoneShorts = []string{
	"ACST",
	"ADT",
	"AEST",
	"AKDT",
	"AST",
	"AWST",
	"BST",
	"CDT",
	"CEST",
	"CST",
	"ChST",
	"EDT",
	"EEST",
	"EET",
	"EST",
	"GMT",
	"GMT+1",
	"GMT+10",
	"GMT+12",
	"GMT+2",
	"GMT+3",
	"GMT+4",
	"GMT+4:30",
	"GMT+5",
	"GMT+5:30",
	"GMT+6",
	"GMT+6:30",
	"GMT+7",
	"GMT+8",
	"GMT+9",
	"GMT+9:30",
	"GMT-3",
	"GMT-4",
	"GMT-5",
	"GMT-6",
	"GMT-7",
	"HKT",
	"HST",
	"IST",
	"JST",
	"MDT",
	"MST",
	"MYT",
	"NDT",
	"NZST",
	"PDT",
	"SGT",
	"UTC−4",
	"WAT",
	"WEST",
}

func TestAbbreviations(t *testing.T) {
	for _, abbr := range timezoneShorts {
		_, err := GetTimezones(abbr)
		if err != nil {
			t.Errorf("Error loading short timezone %s: %v", abbr, err)
			continue
		}
	}
}
