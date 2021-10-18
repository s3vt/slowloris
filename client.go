package slowloris

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var clt *http.Client

func init() {
	transport := &http.Transport{
		MaxConnsPerHost: 0}

	clt = &http.Client{Transport: transport}
}

func DoRequests(requestor Requestor) (string, error) {

	res, err := performRequest(requestor.CreateRequest())
	if err != nil {
		return "", err
	}

	val, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("read response failed %v", err)
		return "", err
	}
	defer res.Body.Close()

	return string(val), nil
}

func performRequest(request *http.Request) (*http.Response, error) {
	res, err := clt.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send request to url %s failed %v", request.URL, err)
	}

	return res, nil
}
