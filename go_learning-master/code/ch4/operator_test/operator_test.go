package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	// &^ 按位置零，d &^ e , 0111 &^ 0001，将前面 把后面对应位置的1变成0， 最终是0110
	t.Log(Readable, Writable, Executable)
	t.Log(5 &^ 0)
	a := 7 //0111
	b := a &^ Readable
	c := a &^ Executable
	t.Log(a, b, c)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
//0111 &^ 0001


