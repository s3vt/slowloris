package slowloris

import "net/http"

type Requestor interface {
	CreateRequest() *http.Request
}
