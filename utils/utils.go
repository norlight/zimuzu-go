package utils

import (
	"fmt"
	"time"
)

func FormatUnix(unix int64) (s string) {
	interval := time.Now().Unix() - unix
	switch {
	case interval < 60:
		return fmt.Sprintf("%d秒前", interval)
	case interval/60 < 60:
		return fmt.Sprintf("%d分钟前", interval/60)
	case interval/(60*60) < 24:
		return fmt.Sprintf("%d小时前", interval/(60*60))
	case interval/(60*60*24) < 30:
		return fmt.Sprintf("%d天前", interval/(60*60*24))
	default:
		t := time.Unix(unix, 0)
		return t.Format("2006-01-02 15:04:05")
	}
}
