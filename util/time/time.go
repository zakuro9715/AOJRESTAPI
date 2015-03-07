package time

import (
	gotime "time"
)

var (
  JST = gotime.FixedZone("Asia/Tokyo", 9 * 60 * 60)
)

func FromUnixTime(d *gotime.Duration, loc *gotime.Location) gotime.Time {
	return gotime.Date(1970, 1, 1, 0, 0, 0, 0, loc).Add(*d)
}
