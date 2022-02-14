package mock

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/golang/mock/gomock"
)

// db mock测试

func TestMockOrderDBI_GetName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockOrderDBI(ctrl)
	mock.EXPECT().GetName(gomock.Eq("teststr"), gomock.Eq(123456)).Return("123456")
	mock.EXPECT().GetName(gomock.Eq("test"), gomock.Eq(1234)).Return("1234")
	mock.EXPECT().GetName(gomock.Eq("test0"), gomock.Eq(0)).Return("0")
	mock.EXPECT().GetName(gomock.Any(), gomock.Eq(999)).Return("any is ok")

	convey.Convey("GetName", t, func() {
		convey.So(mock.GetName("teststr", 123456), convey.ShouldEqual, "123456")
	})

	convey.Convey("GetName", t, func() {
		convey.So(mock.GetName("test", 1234), convey.ShouldEqual, "1234")
	})

	convey.Convey("GetName", t, func() {
		convey.So(mock.GetName("test0", 0), convey.ShouldEqual, "0")
	})

	convey.Convey("GetName", t, func() {
		convey.So(mock.GetName("any", 999), convey.ShouldEqual, "any is ok")
	})
}
