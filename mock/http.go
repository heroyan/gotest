package mock

import "net/http"

func HttpGetWithTimeOut(api string, header map[string]string, timeout int)  {
	http.Get(api)
}
