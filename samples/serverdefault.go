package samples

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type DefaultServer struct {
	*BaseServer
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
