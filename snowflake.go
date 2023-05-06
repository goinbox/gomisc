package gomisc

import (
	"fmt"
	"sync"
	"time"
)

// https://github.com/twitter-archive/snowflake/tree/snowflake-2010
// +----------------------------------------------------------------------------------+
// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit Machine ID  |   12 Bit Sequence Number |
// +----------------------------------------------------------------------------------+

const (
	timestampBits   uint = 41
	machineIDBits   uint = 10
	sequenceNumBits uint = 12

	timestampMax    int64 = -1 ^ (-1 << timestampBits)
	machineIDMax    int64 = -1 ^ (-1 << machineIDBits)
	sequenceNumMask int64 = -1 ^ (-1 << sequenceNumBits)

	timestampShift = sequenceNumBits + machineIDBits

	epochMilli int64 = 1683331200000 // "2023-05-06 00:00:00 UTC"
)

type Snowflake struct {
	lock                sync.Mutex
	lastMilli           int64
	sequenceNum         int64
	machineIDAfterShift int64
}

func NewSnowflake(machineID int64) *Snowflake {
	if machineID > machineIDMax {
		panic(fmt.Sprintf("MachineID %d greater than MachineIDMax %d", machineID, machineIDMax))
	}

	return &Snowflake{
		machineIDAfterShift: machineID << sequenceNumBits,
	}
}

func (s *Snowflake) GenerateID() int64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	milli := time.Now().UnixMilli()
	nowMilli := milli - epochMilli
	if nowMilli > timestampMax {
		panic(fmt.Sprintf("timestamp overflow"))
	}

	if nowMilli == s.lastMilli {
		s.sequenceNum = (s.sequenceNum + 1) & sequenceNumMask
		if s.sequenceNum == 0 {
			for nowMilli <= s.lastMilli {
				nowMilli = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequenceNum = 0
	}

	s.lastMilli = nowMilli

	return nowMilli<<timestampShift | s.machineIDAfterShift | s.sequenceNum
}
