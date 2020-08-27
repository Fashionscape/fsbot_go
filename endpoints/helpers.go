package endpoints

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Requester interface {
	BuildURI() *url.URL
}

func makeRequest(uri *url.URL) ([]byte, error) {
	resp, err := http.Get(uri.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
