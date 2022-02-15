package mock

import (
	"io/ioutil"
	"net/http"
)

func HttpGetWithTimeOut(api string, header map[string]string, timeout int) (string, error) {
	res, err := http.Get(api)
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
