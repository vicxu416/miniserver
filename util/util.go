package util

import "time"

func TimeNow() string {
	return time.Now().Format("2006-01-02T15:04:05")
}
