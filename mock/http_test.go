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
				writer.Write([]byte("hi success"))
				if r.Method != "GET" {
					t.Errorf("Except 'Get' got '%s'", r.Method)
				}
				if r.URL.EscapedPath() != "/" {
					t.Errorf("Expected request to '/', got '%s'", r.URL.EscapedPath())
				}
			}))
			api := ts.URL
			defer ts.Close()
			var header = make(map[string]string)
			HttpGetWithTimeOut(api, header, 30)
			HttpGetWithTimeOut("heiheiheihei", header, 30)
		})
	})
}
