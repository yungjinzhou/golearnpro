package client

import (
	"golearnpro/go_learning-master/code/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	t.Log(series.Square(5))
}
