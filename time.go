/**
* @file misc.go
* @brief misc supermarket
* @author ligang
* @date 2015-12-11
 */

package gomisc

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TimeFmtStrYear   = "2006"
	TimeFmtStrMonth  = "01"
	TimeFmtStrDay    = "02"
	TimeFmtStrHour   = "15"
	TimeFmtStrMinute = "04"
	TimeFmtStrSecond = "05"
)

func TimeGeneralLayout() string {
	return fmt.Sprintf("%s-%s-%s %s:%s:%s",
		TimeFmtStrYear, TimeFmtStrMonth, TimeFmtStrDay, TimeFmtStrHour, TimeFmtStrMinute, TimeFmtStrSecond)
}

func RandByTime(t *time.Time) int64 {
	var timeInt int64

	if t == nil {
		timeInt = time.Now().UnixNano()
	} else {
		timeInt = t.UnixNano()
	}

	return rand.New(rand.NewSource(timeInt)).Int63()
}
