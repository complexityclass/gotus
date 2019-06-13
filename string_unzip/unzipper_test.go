package string_unzip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegularStrings(t *testing.T) {
	//given
	var str1 = "a4bc2d5e"
	var str2 = "abcd"
	var str3 = "45"

	//when
	var res1, err1 = Unzip(str1)
	var res2, err2 = Unzip(str2)
	var _, err3 = Unzip(str3)

	//then
	assert.Equal(t, "aaaabccddddde", res1)
	assert.NoError(t, err1)
	assert.Equal(t, "abcd", res2)
	assert.NoError(t, err2)
	assert.Error(t, err3)
}

func TestEscapedStrings(t *testing.T) {
	//given
	var str1 = "qwe\\4\\5"
	var str2 = "qwe\\45"

	//when
	var res1, err1 = Unzip(str1)
	var res2, err2 = Unzip(str2)

	//then
	assert.Equal(t, "qwe45", res1)
	assert.NoError(t, err1)
	assert.Equal(t, "qwe44444", res2)
	assert.NoError(t, err2)
}

func TestRepeats(t *testing.T) {
	//given
	var str = "a4bc2d5e"
	expected := []Pair{
		{'a', 4},
		{'b', 1},
		{'c', 2},
		{'d', 5},
		{'e', 1}}

	//when
	var reps = repeats(str)

	//then
	assert.Equal(t, expected, reps)
}

func TestDecodeInt(t *testing.T) {
	//given
	var numStr1 = "14578"
	var numStr2 = "124g"

	//when
	var res1, size1 = decodeInt([]byte(numStr1))
	var res2, size2 = decodeInt([]byte(numStr2))

	//then
	assert.Equal(t, 14578, res1)
	assert.Equal(t, 124, res2)
	assert.Equal(t, 5, size1)
	assert.Equal(t, 3, size2)
}
