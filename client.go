package slowloris

import (
	"io"
	"log"
	"net/http"
)

var clt *http.Client

func init() {
	clt = http.DefaultClient
}

func DoRequest(requestor Requestor) (string, error) {
	res, err := clt.Do(requestor.CreateRequest())
	if err != nil {
		log.Printf("send request failed %v", err)
		return "", err
	}

	defer res.Body.Close()

	val, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("read response failed %v", err)
		return "", err
	}

	return string(val), nil
}
