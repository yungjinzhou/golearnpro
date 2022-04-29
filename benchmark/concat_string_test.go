package benchmark

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)


// mac下   go test -bench=.  //  -bench=.代表所有
// -bench=<相关benchmark测试>
// windows 下使用go test 命令行时，-bench=.应写为-bench="."



func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}


func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}


func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0;i < b.N; i++ {
		ret := ""
	for _, elem := range elems {
		ret += elem
	}
	}
	b.StopTimer()
}


func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0;i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}