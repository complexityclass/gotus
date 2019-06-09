package string_unzip

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pair struct {
	a, b interface{}
}

func Unzip(str string) (string, error) {
	var firstRune, _ = utf8.DecodeRune([]byte(str))
	if unicode.IsDigit(firstRune) {
		return "", errors.New("invalid input: string starts with digit")
	}

	var buffer bytes.Buffer
	for _, pair := range repeats(str) {
		var r = pair.a.(rune)
		var c = pair.b.(int)
		var str = strings.Repeat(string(r), c)
		buffer.WriteString(str)
	}

	return buffer.String(), nil
}

func repeats(str string) []Pair {
	var byteStr = []byte(str)
	var reps []Pair
	for rCount := utf8.RuneCount(byteStr); rCount > 0; {
		if rCount == 1 {
			r, _ := utf8.DecodeRune(byteStr)
			reps = append(reps, Pair{r, 1})
			return reps
		}

		fst, fstSize := utf8.DecodeRune(byteStr)
		byteStr = byteStr[fstSize:]

		if fst == '\\' {
			rCount = utf8.RuneCount(byteStr)
			continue
		}

		snd, _ := utf8.DecodeRune(byteStr)

		if !unicode.IsDigit(snd) {
			reps = append(reps, Pair{fst, 1})
			rCount = utf8.RuneCount(byteStr)
			continue
		}

		val, valSize := decodeInt(byteStr)
		byteStr = byteStr[valSize:]
		reps = append(reps, Pair{fst, val})
		rCount = utf8.RuneCount(byteStr)
	}

	return reps
}

func decodeInt(sl []byte) (val int, size int) {
	var tSize = 0
	var rawSlice = sl
	for r, size := utf8.DecodeRune(rawSlice); utf8.RuneCount(rawSlice) > 0; {
		if !unicode.IsDigit(r) {
			break
		} else {
			tSize += size
			rawSlice = rawSlice[size:]
			r, size = utf8.DecodeRune(rawSlice)
		}
	}

	var numSlice = sl[:tSize]
	var value, err = strconv.Atoi(string(numSlice))
	if err == nil {
		return value, tSize
	} else {
		return 0, 0
	}
}
