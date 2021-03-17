# go-timezone

[![GoDocWidget]][GoDocReference] ![Test](https://github.com/tkuchiki/go-timezone/workflows/Test/badge.svg)

[GoDocWidget]:https://godoc.org/github.com/tkuchiki/go-timezone?status.svg
[GoDocReference]:https://godoc.org/github.com/tkuchiki/go-timezone

----

go-timezone is timezone utility for Go.  

It has the following features:

- This library uses only the standard package
- Supports getting offset from timezone abbreviation, which is not supported by the time package
- Determine whether the specified time.Time is daylight saving time
- Change the location of time.Time by specifying the timezone

See [godoc][GoDocReference] for usage.

## Data source

https://github.com/tkuchiki/timezones

# Contributors

- [@alex-tan](https://github.com/alex-tan)
- [@kkavchakHT](https://github.com/kkavchakHT)
- [@scottleedavis](https://github.com/scottleedavis)
- [@sashabaranov](https://github.com/sashabaranov)
