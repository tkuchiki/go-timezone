package timezone

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTzAbbreviationInfo(t *testing.T) {
	t.Parallel()

	tz := New()

	abbr := "EET"
	want := 7200
	t.Run("EET is offset 7200", func(t *testing.T) {
		tzAbbrInfo, err := tz.GetTzAbbreviationInfo(abbr)
		if err != nil {
			t.Fatal(err)
		}

		got := tzAbbrInfo[0].Offset()
		if got != want {
			t.Fatalf(`want: %v, got: %v`, want, got)
		}
	})

	t.Run("BST are ambiguous abbreviations", func(t *testing.T) {
		ambiguousAbbr := "BST"
		_, err := tz.GetTzAbbreviationInfo(ambiguousAbbr)
		if err != ErrAmbiguousTzAbbreviations {
			t.Fatal(err)
		}
	})
}

func TestGetTzAbbreviationInfoByTZName(t *testing.T) {
	t.Parallel()

	tz := New()

	abbr := "BST"
	t.Run("BST / British Summer Time is offset 3600", func(t *testing.T) {
		name := "British Summer Time"
		offset := 3600

		tzAbbrInfo, err := tz.GetTzAbbreviationInfoByTZName(abbr, name)
		if err != nil {
			t.Fatal(err)
		}

		if tzAbbrInfo.Offset() != offset {
			t.Fatalf(`want: %d, got: %d`, offset, tzAbbrInfo.Offset())
		}
	})

	t.Run("Invalid Time", func(t *testing.T) {
		name := "Invalid Time"
		tzAbbrInfo, _ := tz.GetTzAbbreviationInfoByTZName(abbr, name)
		if tzAbbrInfo != nil {
			t.Fatalf(`want: "%T", got: "%T"`, nil, tzAbbrInfo)
		}
	})
}

func TestGetTimezones(t *testing.T) {
	t.Parallel()

	tz := New()

	t.Run("UTC", func(t *testing.T) {
		abbr := "UTC"
		want := []string{
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

		if !reflect.DeepEqual(want, timezones) {
			t.Fatalf(`want: %v, got: %v`, want, timezones)
		}
	})

	t.Run("Invalid/Zone", func(t *testing.T) {
		abbr := "Invalid"
		_, err := tz.GetTimezones(abbr)
		if err == nil {
			t.Fatal("Invalid timezone")
		}
	})
}

func TestFixedTimezone(t *testing.T) {
	t.Parallel()

	tz := New()

	baseTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	t.Run("Europe/Belgrade", func(t *testing.T) {
		timezone := "Europe/Belgrade"
		want := 3600

		_time, err := tz.FixedTimezone(baseTime, timezone)
		if err != nil {
			t.Fatal(err)
		}

		_, offset := _time.Zone()
		if want != offset {
			t.Fatalf(`want: %d, got: %d`, want, offset)
		}
	})

	t.Run("Invalid/Zone", func(t *testing.T) {
		timezone := "Invalid/Zone"
		_time, _ := tz.FixedTimezone(baseTime, timezone)
		if !_time.IsZero() {
			t.Fatal("Invalid timezone")
		}
	})
}

func TestGetTzInfo(t *testing.T) {
	t.Parallel()

	tz := New()

	t.Run("Europe/London", func(t *testing.T) {
		timezone := "Europe/London"
		want := "GMT"
		tzInfo, err := tz.GetTzInfo(timezone)
		if err != nil {
			t.Fatal(err)
		}

		abbr := tzInfo.ShortStandard()
		if abbr != want {
			t.Fatalf(`want: %s, got: %s`, want, abbr)
		}

		want = "BST"
		abbr = tzInfo.ShortDaylight()
		if abbr != want {
			t.Fatalf(`want: %s, got: %s`, want, abbr)
		}

		if !tzInfo.HasDST() {
			t.Fatalf(`want: true, got: %v`, tzInfo.HasDST())
		}
	})

	t.Run("Invalid/Zone", func(t *testing.T) {
		timezone := "Invalid/Zone"
		tzInfo, _ := tz.GetTzInfo(timezone)
		if tzInfo != nil {
			t.Fatal("Invalid timezone")
		}
	})
}

func TestGetOffset(t *testing.T) {
	t.Parallel()

	tz := New()

	t.Run("GMT offset is 0", func(t *testing.T) {
		abbr := "GMT"
		want := 0
		offset, err := tz.GetOffset(abbr)
		if err != nil {
			t.Fatal(err)
		}

		if offset != want {
			t.Fatalf(`want: %d, got: %d`, want, offset)
		}
	})

	t.Run("BST are ambiguous abbreviations", func(t *testing.T) {
		abbr := "BST"
		_, err := tz.GetOffset(abbr, true)
		if err != ErrAmbiguousTzAbbreviations {
			t.Fatal(err)
		}
	})

	t.Run("Invalid abbreviation", func(t *testing.T) {
		abbr := "Invalid"
		offset, _ := tz.GetOffset(abbr)
		if offset != 0 {
			t.Fatal("Invalid timezone abbreviation")
		}
	})

}

func TestGetTimezoneAbbreviation(t *testing.T) {
	t.Parallel()

	tz := New()

	cases := map[string]struct {
		want     string
		dst      bool
		timezone string
		skipErr  bool
	}{
		"Europe/London is GMT": {
			want:     "GMT",
			dst:      false,
			timezone: "Europe/London",
			skipErr:  false,
		},
		"Europe/London is BST when DST": {
			want:     "BST",
			dst:      true,
			timezone: "Europe/London",
			skipErr:  false,
		},
		"Invalid/Zone is empty": {
			want:     "",
			timezone: "Invalid/Zone",
			skipErr:  true,
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			abbr, err := tz.GetTimezoneAbbreviation(tc.timezone, tc.dst)
			if !tc.skipErr && err != nil {
				t.Fatal(err)
			}

			if abbr != tc.want {
				t.Fatalf(`want: %s, got: %s`, tc.want, abbr)
			}
		})
	}
}

func TestIsDST(t *testing.T) {
	t.Parallel()

	tz := New()

	locNY, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}

	locUTC, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatal(err)
	}

	cases := map[string]struct {
		want bool
		loc  *time.Location
		t    time.Time
	}{
		"America/New_York is not DST on 2021-01-01 at 0:00": {
			want: false,
			loc:  locNY,
			t:    time.Date(2021, 1, 1, 1, 0, 0, 0, locNY),
		},
		"America/New_York is DST on 2021-07-01 at 0:00": {
			want: true,
			loc:  locNY,
			t:    time.Date(2021, 7, 1, 1, 0, 0, 0, locNY),
		},
		"UTC is not DST on 2021-01-01 at 0:00": {
			want: false,
			loc:  locUTC,
			t:    time.Date(2021, 1, 1, 1, 0, 0, 0, locUTC),
		},
		"UTC is not DST on 2021-07-01 at 0:00": {
			want: false,
			loc:  locUTC,
			t:    time.Date(2021, 7, 1, 1, 0, 0, 0, locUTC),
		},
	}

	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			isDST := tz.IsDST(tc.t)
			if tc.want != isDST {
				t.Fatalf(`want: %v, got: %v`, tc.want, isDST)
			}
		})
	}
}
