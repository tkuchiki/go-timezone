# go-timezone
Timezone utility for Golang

## Example

### Code

```
package main

import (
	"fmt"
    "github.com/tkuchiki/go-timezone"
)

func main() {
	offset, err := GetOffset("JST")
	fmt.Println(offset, err)

	offset, err = GetOffset("hogehoge")
	fmt.Println(offset, err)

	var zones []string
	zones, err = GetTimezones("UTC")
	fmt.Println(zones, err)

	zones, err = GetTimezones("foobar")
	fmt.Println(zones, err)
}

```

### Result

```
32400 <nil>
0 Invalid short timezone: hogehoge
[Antarctica/Troll Etc/UTC Etc/Universal Etc/Zulu UTC Universal Zulu] <nil>
[] Invalid short timezone: foobar
```
