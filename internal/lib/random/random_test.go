package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRandomString(t *testing.T) {
	testTable := []struct {
		name string
		size int
	}{
		{
			name: "test 1",
			size: 1,
		},
		{
			name: "test 2",
			size: 2,
		},
		{
			name: "test 3",
			size: 5,
		},
		{
			name: "test 4",
			size: 10,
		},
		{
			name: "test 5",
			size: 20,
		},
	}

	for _, testCase := range testTable {
		str1 := NewRandomString(testCase.size)
		str2 := NewRandomString(testCase.size)
		assert.Len(t, str1, testCase.size)
		assert.Len(t, str2, testCase.size)
		assert.NotEqual(t, str1, str2)
	}
}
