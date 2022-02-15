package mock

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
)

func TestNewMyCls(t *testing.T) {
	my := NewMyCls("shuifa", 18)
	convey.Convey("my.GetName", t, func() {
		convey.So(my.GetName(), convey.ShouldEqual, "shuifa")
	})
}

func TestMonkeyFunc(t *testing.T) {
	convey.Convey("MonkeyFunc test", t, func() {
		patches := gomonkey.ApplyFunc(MonkeyFunc, func(a, b int) int {
			return a*b*a*b
		})
		defer patches.Reset()
		c := MonkeyFunc(2, 3)
		convey.So(c, convey.ShouldEqual, 36)
	})

	convey.Convey("MonkeyFuncSeq test", t, func() {
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{1}},
			{Values: gomonkey.Params{2}},
			{Values: gomonkey.Params{3}},
			{Values: gomonkey.Params{4}},
		}
		patches := gomonkey.ApplyFuncSeq(MonkeyFunc, outputs)
		defer patches.Reset()
		convey.So(MonkeyFunc(1, 2), convey.ShouldEqual, 1)
		convey.So(MonkeyFunc(1, 2), convey.ShouldEqual, 2)
		convey.So(MonkeyFunc(1, 2), convey.ShouldEqual, 3)
		convey.So(MonkeyFunc(1, 2), convey.ShouldEqual, 4)
	})
}

func TestGlobalVar(t *testing.T) {
	convey.Convey("GlobalVar test", t, func() {
		convey.Convey("change", func() {
			patches := gomonkey.ApplyGlobalVar(&global1, 888)
			defer patches.Reset()
			convey.So(global1, convey.ShouldEqual, 888)
		})

		convey.Convey("recover", func() {
			convey.So(global1, convey.ShouldEqual, 100)
		})
	})
}

func TestApplyMethod(t *testing.T)  {
	my := NewMyCls("fage", 18)
	var s *MyCls
	convey.Convey("ApplyMethod test", t, func() {
		convey.Convey("for success", func() {
			patches := gomonkey.ApplyMethod(reflect.TypeOf(s), "SayHi", func(_ *MyCls) {
				fmt.Println("hi, i am monkey")
			})
			defer patches.Reset()
			patches.ApplyMethod(reflect.TypeOf(s), "GetName", func(_ *MyCls) string {
				return "fagenb"
			})
			my.SayHi()
			convey.So(my.GetName(), convey.ShouldEqual, "fagenb")
		})
	})
}