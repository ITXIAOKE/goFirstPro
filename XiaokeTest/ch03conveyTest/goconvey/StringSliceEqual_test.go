package goconvey

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringSliceEqual(t *testing.T){
	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey2"}
		So(StringSliceEqual(a, b), ShouldBeTrue)
	})
}