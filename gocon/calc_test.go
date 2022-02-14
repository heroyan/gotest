package gocon

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFindTheBiggest(t *testing.T) {
	convey.Convey("Find the biggest one of the array", t, func() {
		one := FindTheBiggest([]int{1, 2, 3, 44, 33, 22, 343242, 222})
		convey.So(one, convey.ShouldEqual, 343242)
	})

	convey.Convey("Find the biggest one of the array", t, func() {
		one := FindTheBiggest([]int{})
		convey.So(one, convey.ShouldEqual,-1)
	})

	convey.Convey("Find the biggest one of the array", t, func() {
		one := FindTheBiggest([]int{0})
		convey.So(one, convey.ShouldEqual,0)
	})
}

func TestFindIfIn(t *testing.T) {
	convey.Convey("find 2 if in array [1,2,3,4,5,7]", t, func() {
		one := FindIfIn([]int{1,2,3,4,5,7}, 7)
		convey.So(one, convey.ShouldBeTrue)
	})

	convey.Convey("find 12 if in array [1,2,3,4,5,7]", t, func() {
		one := FindIfIn([]int{1,2,3,4,5,7}, 12)
		convey.So(one, convey.ShouldBeFalse)
	})

	convey.Convey("find 12 if in array []", t, func() {
		one := FindIfIn([]int{}, 12)
		convey.So(one, convey.ShouldBeFalse)
	})
}