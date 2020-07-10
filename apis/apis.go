package apis

import (
	"io/ioutil"
	"net/http"
)

var Client httpInterface = &httpClient{}

type httpInterface interface {
	Get(string) ([]byte, error)
}

type httpClient struct{}

func (c *httpClient) Get(url string) ([]byte, error) {
	return httpGet(url)
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(resp.Body)
}
