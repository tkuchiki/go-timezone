# go-timezone
Timezone utility for Golang

## Example

### Code

```go
package main

import (
    "fmt"
    "github.com/tkuchiki/go-timezone"
    "time"
)

func main() {
	offset, err := timezone.GetOffset("JST")
	fmt.Println(offset, err)

	offset, err = timezone.GetOffset("hogehoge")
	fmt.Println(offset, err)

	var zones []string
	zones, err = timezone.GetTimezones("UTC")
	fmt.Println(zones, err)

	zones, err = timezone.GetTimezones("foobar")
	fmt.Println(zones, err)

	now := time.Now()

	fmt.Println("## current timezone")
	fmt.Println(now)

	var jst time.Time
	loc, _ := time.LoadLocation("UTC")
	utc := now.In(loc)

	jst, _ = timezone.FixedTimezone(utc, "")

	fmt.Println("## UTC")
	fmt.Println(utc)
	fmt.Println("## UTC -> current timezone")
	fmt.Println(jst)

	var est time.Time
	est, _ = timezone.FixedTimezone(now, "America/New_York")
	fmt.Println("## current timezone -> EDT")
	fmt.Println(est)

    offset, _ = timezone.GetOffset("EST")
    zone, _ := timezone.GetTimezoneAbbreviation("America/New_York")
    fmt.Println("## DST is not considered")
    fmt.Println(offset)
    fmt.Println(zone)

    offset, _ = timezone.GetOffset("EST", true)
    zone, _ = timezone.GetTimezoneAbbreviation("America/New_York", true)
    fmt.Println("## DST is considered")
    fmt.Println(offset)
    fmt.Println(zone)
}
```

### Result

```console
# current timezone = UTC
$ TZ=UTC go run /path/to/main.go
32400 <nil>
0 Invalid short timezone: hogehoge
[Antarctica/Troll Etc/UTC Etc/Universal Etc/Zulu UTC Universal Zulu] <nil>
[] Invalid short timezone: foobar
## current timezone
2018-03-15 00:07:01.921041165 +0000 UTC m=+0.000669042
## UTC
2018-03-15 00:07:01.921041165 +0000 UTC
## UTC -> current timezone
2018-03-15 00:07:01.921041165 +0000 UTC
## current timezone -> EDT
2018-03-14 20:07:01.921041165 -0400 EDT
## DST is not considered
-18000
EST
## DST is considered
-14400
EDT

# current timezone = JST
$ TZ=Asia/Tokyo go run /path/to/main.go
32400 <nil>
0 Invalid short timezone: hogehoge
[Antarctica/Troll Etc/UTC Etc/Universal Etc/Zulu UTC Universal Zulu] <nil>
[] Invalid short timezone: foobar
## current timezone
2018-03-15 09:08:58.410680998 +0900 JST m=+0.000536318
## UTC
2018-03-15 00:08:58.410680998 +0000 UTC
## UTC -> current timezone
2018-03-15 09:08:58.410680998 +0900 JST
## current timezone -> EDT
2018-03-14 20:08:58.410680998 -0400 EDT
## DST is not considered
-18000
EST
## DST is considered
-14400
EDT
```

# Contributors

- [@alex-tan](https://github.com/alex-tan)
- [@kkavchakHT](https://github.com/kkavchakHT)
- [@scottleedavis](https://github.com/scottleedavis)
- [@sashabaranov](https://github.com/sashabaranov)
