package problem1003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans bool
}{

	{
		"aabcbc",
		true,
	},

	{
		"abcabcababcc",
		true,
	},

	{
		"abccba",
		false,
	},

	{
		"cababc",
		false,
	},

	// 可以有多个 testcase
}

func Test_isValid(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isValid(tc.S), "输入:%v", tc)
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isValid(tc.S)
		}
	}
}
