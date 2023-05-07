package utils

import "time"

func MakeTimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
