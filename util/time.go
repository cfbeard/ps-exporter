package util

import "time"

func MsToTime(ms int64) time.Time {
    return time.Unix(0, ms * int64(time.Millisecond))
}
