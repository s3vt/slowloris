package slowloris

import (
	"fmt"
	"io"
	"net/http"
)

type Requestor interface {
	CreateRequest() *http.Request
}

type BaseRequestor struct {
	Method, Scheme,
	Host, Port, Path string
	Body io.Reader
}

func (b *BaseRequestor) Url() string {
	return fmt.Sprintf("%s://%s:%s/%s", b.Scheme, b.Host, b.Port, b.Path)
}
