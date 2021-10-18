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
		Username: samples.BasAuthUsr,
		Password: samples.BasAuthPwd,

		BaseRequestor: &slowloris.BaseRequestor{
			Scheme: "http",
			Method: http.MethodPost,
			Host:   samples.Host,
			Port:   samples.Port,
			Path:   "",
			Body: &slowloris.SlowReader{
				Content:   "hellotherekjoksajfn roqibncksanclsakdj csahfdslakdsdkcnajsrjvjaskjhdfsl csdajpj",
				SleepTime: 500 * time.Millisecond,
				Log:       true}}}

	res, err := slowloris.DoRequests(requestor)
	if err != nil {
		log.Fatalf("error in request %v", err)
	}
	log.Printf("received response %s", res)
}
