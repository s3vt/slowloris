package samples

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sapvs/slowloris"
)

type DefaultServer struct {
	*BaseServer
}

type DefaultRequestor struct {
	*slowloris.BaseRequestor
}

func (requestor *DefaultRequestor) CreateRequest() *http.Request {
	req, err := http.NewRequestWithContext(context.Background(), requestor.Method, fmt.Sprintf("http://%s:%s/%s", requestor.Host, requestor.Port, requestor.Path), requestor.Body)
	if err != nil {
		log.Fatalf("could not create basic auth request due to %v", err)
	}

	return req
}

func (ds *DefaultServer) Start() error {
	http.HandleFunc("/", root)
	return http.ListenAndServe(ds.Address, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("could not read body due to %v", err)))
		return
	}
	defer r.Body.Close()

	log.Print(string(body))

	w.Write(body)
}
