package bbd

import (
	"testing"
         . "github.com/smartstreets/goconvey/convey"
)


//项目网站： https://github.com/smartstreets/goconvey
// 安装： go get -u github.com/smartstreets/goconvey/convey
// 启动web ui
// $GOPATH/bin/goconvey


func TestSpec(t *testing.T) {
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4
		Convey("When add the two numbers", func() {
			c := a + b
			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}


