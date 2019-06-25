package otus_logger

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccepted(t *testing.T) {
	//given
	var accepted = HwAccepted{24345, 5}
	var buf bytes.Buffer
	var expected = fmt.Sprintf("%s accepted 24345 5", time.Now().Format("01-02-2006"))

	//when
	LogOtusEvent(accepted, &buf)

	//then
	assert.Equal(t, expected, buf.String())
}

func TestSubmitted(t *testing.T) {
	//given
	var accepted = HwSubmitted{24345, "404", "some comment"}
	var buf bytes.Buffer
	var expected = fmt.Sprintf("%s submitted 24345 some comment", time.Now().Format("01-02-2006"))

	//when
	LogOtusEvent(accepted, &buf)

	//then
	assert.Equal(t, expected, buf.String())
}
