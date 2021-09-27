package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sapvs/slowloris"
	"github.com/sapvs/slowloris/samples"
)

func main() {
	submitbasicauthrequest()
}

func submitbasicauthrequest() {
	requestor := &samples.BasicAuthRequestor{
		Method:   http.MethodPost,
		Host:     samples.Host,
		Port:     samples.Port,
		Path:     samples.Path,
		Username: samples.BasicAuthUser,
		Password: samples.BasicAuthPassword,
		Body:     &slowloris.SlowReader{Content: "hello", SleepTime: 500 * time.Millisecond, Log: true}}

	res, err := slowloris.DoRequest(requestor)
	if err != nil {
		log.Fatalf("error in request %v", err)
	}
	log.Printf("received response %s", res)
}
