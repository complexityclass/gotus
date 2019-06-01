package mtimelog

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

const (
	ntpHost = "0.ru.pool.ntp.org"
)

func LogCurrentTime() (string, error) {
	var time, err = ntp.Time(ntpHost)
	if err == nil {
		return time.String(), nil
	} else {
		return "", err
	}
}

func HelloNow() {
	var timeString, err = LogCurrentTime()
	if err != nil {
		log.Printf("can't obtain time with %s", err)
	} else {
		fmt.Printf("current time is %s", timeString)
	}
}
