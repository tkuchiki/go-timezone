package timezone

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTzAbbreviationInfo(t *testing.T) {
	t.Parallel()

	tz := New()

	test := struct {
		abbr      string
		offset    int
		errOffset string
	}{
		abbr:      "EET",
		offset:    7200,
		errOffset: `expected: 7200, actual: %d`,
	}

	tzAbbrInfo, err := tz.GetTzAbbreviationInfo(test.abbr)
	if err != nil {
		t.Fatal(err)
	}

	if tzAbbrInfo[0].Offset() != test.offset {
		t.Fatalf(test.errOffset, test.offset)
	}

	ambiguousAbbr := "BST"
	tzAbbrInfo, err = tz.GetTzAbbreviationInfo(ambiguousAbbr)
	if err != ErrAmbiguousTzAbbreviations {
		t.Fatal(err)
	}
}

func TestGetTzAbbreviationInfoByTZName(t *testing.T) {
	t.Parallel()

	tz := New()

	abbr := "BST"
	name := "British Summer Time"
	offset := 3600

	tzAbbrInfo, err := tz.GetTzAbbreviationInfoByTZName(abbr, name)
	if err != nil {
		t.Fatal(err)
	}

	if tzAbbrInfo.Offset() != offset {
		t.Fatalf(`expected: %d, actual: %d`, offset, tzAbbrInfo.Offset())
	}

	name = "Invalid Time"
	tzAbbrInfo, _ = tz.GetTzAbbreviationInfoByTZName(abbr, name)
	if tzAbbrInfo != nil {
		t.Fatalf(`expected: "%T", actual: "%T"`, nil, tzAbbrInfo)
	}
}

func TestGetTimezones(t *testing.T) {
	t.Parallel()

	tz := New()

	abbr := "UTC"
	expectedTimezones := []string{
		"Etc/UCT",
		"Etc/UTC",
		"Etc/Universal",
		"Etc/Zulu",
		"UCT",
		"UTC",
		"Universal",
		"Zulu",
	}

	timezones, err := tz.GetTimezones(abbr)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expectedTimezones, timezones) {
		t.Fatalf(`expected: %v, actual: %v`, expectedTimezones, timezones)
	}

	abbr = "Invalid"
	_, err = tz.GetTimezones(abbr)
	if err == nil {
		t.Fatal("Invalid timezone")
	}
}

func TestFixedTimezone(t *testing.T) {
	t.Parallel()

	tz := New()

	baseTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	timezone := "Europe/Belgrade"
	expect := 3600

	_time, err := tz.FixedTimezone(baseTime, timezone)
	if err != nil {
		t.Fatal(err)
	}

	_, offset := _time.Zone()
	if expect != offset {
		t.Fatalf(`expected: %d, actual: %d`, expect, offset)
	}

	timezone = "Invalid/Zone"
	_time, _ = tz.FixedTimezone(baseTime, timezone)
	if !_time.IsZero() {
		t.Fatal("Invalid timezone")
	}
}

func TestGetTzInfo(t *testing.T) {
	t.Parallel()

	tz := New()

	timezone := "Europe/London"
	expect := "GMT"
	tzInfo, err := tz.GetTzInfo(timezone)
	if err != nil {
		t.Fatal(err)
	}

	abbr := tzInfo.ShortStandard()
	if abbr != expect {
		t.Fatalf(`expected: %s, actual: %s`, expect, abbr)
	}

	expect = "BST"
	abbr = tzInfo.ShortDaylight()
	if abbr != expect {
		t.Fatalf(`expected: %s, actual: %s`, expect, abbr)
	}

	if !tzInfo.HasDST() {
		t.Fatalf(`expected: true, actual: %v`, tzInfo.HasDST())
	}

	timezone = "Invalid/Zone"
	tzInfo, _ = tz.GetTzInfo(timezone)
	if tzInfo != nil {
		t.Fatal("Invalid timezone")
	}
}

func TestGetOffset(t *testing.T) {
	t.Parallel()

	tz := New()

	abbr := "GMT"
	expect := 0
	offset, err := tz.GetOffset(abbr)
	if err != nil {
		t.Fatal(err)
	}

	if offset != expect {
		t.Fatalf(`expected: %d, actual: %d`, expect, offset)
	}

	abbr = "BST"
	_, err = tz.GetOffset(abbr, true)
	if err != ErrAmbiguousTzAbbreviations {
		t.Fatal(err)
	}

	abbr = "Invalid"
	offset, _ = tz.GetOffset(abbr)
	if offset != 0 {
		t.Fatal("Invalid timezone abbreviation")
	}
}

func TestGetTimezoneAbbreviation(t *testing.T) {
	t.Parallel()

	tz := New()

	timezone := "Europe/London"
	expect := "GMT"
	abbr, err := tz.GetTimezoneAbbreviation(timezone)
	if err != nil {
		t.Fatal(err)
	}

	if abbr != expect {
		t.Fatalf(`expected: %s, actual: %s`, expect, abbr)
	}

	expect = "BST"
	abbr, err = tz.GetTimezoneAbbreviation(timezone, true)
	if err != nil {
		t.Fatal(err)
	}

	if abbr != expect {
		t.Fatalf(`expected: %s, actual: %s`, expect, abbr)
	}

	timezone = "Invalid/Zone"
	abbr, _ = tz.GetTimezoneAbbreviation(timezone)
	if abbr != "" {
		t.Fatal("Invalid timezone")
	}
}
