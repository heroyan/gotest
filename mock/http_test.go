package mock

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestHttpGetWithTimeOut(t *testing.T) {
	convey.Convey("TestHttp", t, func() {
		convey.Convey("test normal", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
				writer.WriteHeader(http.StatusOK)
				if r.Method != "GET" {
					t.Errorf("Except 'Get' got '%s'", r.Method)
				}
				if r.URL.EscapedPath() == "/get/articles" {
					writer.Write([]byte("100 articles"))
				}
				if r.URL.EscapedPath() == "/get/notices" {
					writer.Write([]byte("1000 notices"))
				}
			}))
			api := ts.URL
			t.Log(api)
			defer ts.Close()
			var header = make(map[string]string)
			content, err := HttpGetWithTimeOut(api + "/get/articles", header, 30)
			convey.So(err, convey.ShouldBeNil)
			convey.So(content, convey.ShouldEqual, "100 articles")
		})
	})
}
