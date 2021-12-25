package internal

import (
	"net/http"

	"github.com/matthewzhaocc/planetary-deploy/internal/org"
	"github.com/matthewzhaocc/planetary-deploy/internal/types"
	"github.com/matthewzhaocc/planetary-deploy/internal/user"
)

func GetMux(serv types.Server) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/org", org.AddOrgRoute(serv))
	mux.HandleFunc("/user/", user.AddUserRoute(serv))
	mux.HandleFunc("/auth/login", user.Authenticate(serv))
	return mux
}
