package otus_logger

import (
	"fmt"
	"io"
	"log"
	"time"
)

type TimeProducer interface {
	GetFormattedTime() string
}

type OtusEvent interface {
	fmt.Stringer
	TimeProducer
}

type HwAccepted struct {
	Id    int
	Grade int
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

func (HwAccepted) GetFormattedTime() string {
	return time.Now().Format("01-02-2006")
}

func (HwSubmitted) GetFormattedTime() string {
	return time.Now().Format("01-02-2006")
}

func (status HwAccepted) String() string {
	return fmt.Sprintf("%s accepted %d %d", status.GetFormattedTime(), status.Id, status.Grade)
}

func (status HwSubmitted) String() string {
	return fmt.Sprintf("%s submitted %d %s", status.GetFormattedTime(), status.Id, status.Comment)
}

func LogOtusEvent(event OtusEvent, writer io.Writer) {
	_, err := writer.Write([]byte(event.String()))
	if err != nil {
		log.Fatal(err)
	}
}
