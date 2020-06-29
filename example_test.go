package timezone_test

import (
	"fmt"
	"time"

	"github.com/tkuchiki/go-timezone"
)

func ExampleTimezone_GetTzAbbreviationInfo() {
	tz := timezone.New()
	tzAbbrInfos, _ := tz.GetTzAbbreviationInfo("EET")

	fmt.Println(tzAbbrInfos[0].Name())
	fmt.Println(tzAbbrInfos[0].Offset())
	fmt.Println(tzAbbrInfos[0].OffsetHHMM())
	// Output:
	// Eastern European Time/Eastern European Standard Time
	// 7200
	// +02:00
}

func ExampleTimezone_GetTzAbbreviationInfo_ambiguousTimezoneAbbreviationsError() {
	tz := timezone.New()
	tzAbbrInfos, _ := tz.GetTzAbbreviationInfo("BST")

	fmt.Println(tzAbbrInfos[0].Name())
	fmt.Println(tzAbbrInfos[0].Offset())
	fmt.Println(tzAbbrInfos[0].OffsetHHMM())
	fmt.Println(tzAbbrInfos[1].Name())
	fmt.Println(tzAbbrInfos[1].Offset())
	fmt.Println(tzAbbrInfos[1].OffsetHHMM())
	fmt.Println(tzAbbrInfos[2].Name())
	fmt.Println(tzAbbrInfos[2].Offset())
	fmt.Println(tzAbbrInfos[2].OffsetHHMM())
	// Output:
	// Bolivia Summer Time
	// -12756
	// -03:27
	// British Summer Time
	// 3600
	// +01:00
	// Bougainville Standard Time
	// 39600
	// +11:00
}

func ExampleTimezone_GetTzAbbreviationInfoByTZName() {
	tz := timezone.New()
	tzAbbrInfo, _ := tz.GetTzAbbreviationInfoByTZName("BST", "British Summer Time")

	fmt.Println(tzAbbrInfo.Name())
	fmt.Println(tzAbbrInfo.Offset())
	fmt.Println(tzAbbrInfo.OffsetHHMM())
	// Output:
	// British Summer Time
	// 3600
	// +01:00
}

func ExampleTimezone_GetTimezones() {
	tz := timezone.New()
	timezones, _ := tz.GetTimezones("UTC")

	fmt.Println(timezones)
	// Output:
	// [Etc/UCT Etc/UTC Etc/Universal Etc/Zulu UCT UTC Universal Zulu]
}

func ExampleTimezone_FixedTimezone() {
	tz := timezone.New()

	_time := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fixedTime, _ := tz.FixedTimezone(_time, "Europe/Belgrade")

	fmt.Println(fixedTime)
	// Output:
	// 2020-01-01 01:00:00 +0100 CET
}

func ExampleTimezone_GetTzInfo() {
	tz := timezone.New()
	tzInfo, _ := tz.GetTzInfo("Europe/London")

	fmt.Println(tzInfo.LongStandard())
	fmt.Println(tzInfo.ShortStandard())
	fmt.Println(tzInfo.StandardOffset())
	fmt.Println(tzInfo.StandardOffsetHHMM())
	fmt.Println(tzInfo.LongDaylight())
	fmt.Println(tzInfo.ShortDaylight())
	fmt.Println(tzInfo.DaylightOffset())
	fmt.Println(tzInfo.DaylightOffsetHHMM())
	// Output:
	// Greenwich Mean Time
	// GMT
	// 0
	// +00:00
	// British Summer Time
	// BST
	// 3600
	// +01:00
}

func ExampleTimezone_GetTimezoneAbbreviation() {
	tz := timezone.New()

	abbr, _ := tz.GetTimezoneAbbreviation("Europe/London")

	fmt.Println(abbr)
	// Output:
	// GMT
}

func ExampleTimezone_GetTimezoneAbbreviation_dst() {
	tz := timezone.New()

	abbr, _ := tz.GetTimezoneAbbreviation("Europe/London", true)

	fmt.Println(abbr)
	// Output:
	// BST
}
