package gotus

import (
	"fmt"
	"github.com/beevik/ntp"
)

const (
	ntpHost = "0.ru.pool.ntp.org"
	UIError = "server unreachable"
)

func LogCurrentTime() string {
	var time, err = ntp.Time(ntpHost)
	if err == nil {
		return time.String()
	} else {
		return UIError
	}
}

func HelloNow() {
	fmt.Println(LogCurrentTime())
}
