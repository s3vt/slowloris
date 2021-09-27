package samples

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	usr, pwd string
)

type BasicAuthRequestor struct {
	Method, Host, Port, Path, Username, Password string
	Body                                         io.Reader
}

func (requestor *BasicAuthRequestor) CreateRequest() *http.Request {
	req, err := http.NewRequestWithContext(context.Background(), requestor.Method, fmt.Sprintf("http://%s:%s/%s", requestor.Host, requestor.Port, requestor.Path), requestor.Body)
	if err != nil {
		log.Fatalf("could not create basic auth request due to %v", err)
	}

	req.SetBasicAuth(requestor.Username, requestor.Password)

	return req
}

type BasicAuthServer struct {
	*BaseServer
	Username, Password string
}

func (bas *BasicAuthServer) Start() error {
	usr, pwd = bas.Username, bas.Password

	http.HandleFunc("hello", authhandler)
	return http.ListenAndServe(bas.Address, nil)
}

func authhandler(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect username password"))
		return
	}

	if u != usr && p != pwd {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect username password"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("could not read body due to %v", err)))
		return
	}
	defer r.Body.Close()

	log.Printf("Body %s", body)
	w.Write([]byte(body))
}
