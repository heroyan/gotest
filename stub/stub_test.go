package stub

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/prashantv/gostub"
)

func TestShowAndGetMyName(t *testing.T) {
	stubs := gostub.Stub(&myName, "fage nb")
	defer stubs.Reset()

	convey.Convey("fage nb", t, func() {
		convey.So(ShowAndGetMyName(), convey.ShouldEqual, "fage nb")
	})

	stubs.Stub(&myName, "ks nb")
	convey.Convey("ks nb", t, func() {
		convey.So(ShowAndGetMyName(), convey.ShouldEqual, "ks nb")
	})
}

func TestShowAndGetMyName2(t *testing.T) {
	convey.Convey("yanshuifa", t, func() {
		convey.So(ShowAndGetMyName(), convey.ShouldEqual, "yanshuifa")
	})
}

func TestShowAndGetMyName3(t *testing.T) {
	stubs := gostub.Stub(&show, func(name string) string {
		fmt.Println("in stub my name")

		return "solid value"
	})
	defer stubs.Reset()

	convey.Convey("yanshuifa", t, func() {
		convey.So(show("shuifa"), convey.ShouldEqual, "solid value")
	})
}

func TestShowAndGetMyName4(t *testing.T) {
	convey.Convey("yanshuifa", t, func() {
		convey.So(show("shuifa"), convey.ShouldEqual, "shuifa")
	})
}

func TestShowAndGetMyName5(t *testing.T) {
	stubs := gostub.StubFunc(&show, "niu niu niu")
	defer stubs.Reset()

	convey.Convey("StubFunc", t, func() {
		convey.So(show("shuifa"), convey.ShouldEqual, "niu niu niu")
	})
}
