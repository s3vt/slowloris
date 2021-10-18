package samples

const (
	Host    = "localhost"
	Port    = "8080"
	Address = Host + ":" + Port
	Path    = "path"

	BasAuthUsr = "user"
	BasAuthPwd = "password"
)

type BaseServer struct {
	Address string
}

type Server interface {
	Start() error
}
