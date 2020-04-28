package utils

import (
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"
)

//截取指定字符串之间的字符
func SubstringBetween(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := -1
	for i := posFirst + len(a); i < len(value); i++ {
		s := value[i:(i + len(b))]
		if s == b {
			posLast = i
			break
		}
	}
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

var qnLock sync.Mutex

//生成QN
func GenerateQNStr() string {
	qnLock.Lock()
	defer func() {
		time.Sleep(time.Microsecond * 100)
		qnLock.Unlock()
	}()
	now := time.Now().UnixNano()
	miliSeconds := (now % 1e9) / 1e6
	return fmt.Sprintf("%04d%02d%02d%02d%02d%02d%03d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), miliSeconds)
}

//lrc检验码
func GenerateLRC(data []byte) string {
	sum := uint8(0)
	for _, b := range data {
		sum += b
	}
	u := uint8(-int8(sum))
	return strings.ToUpper(hex.EncodeToString([]byte{u}))
}
