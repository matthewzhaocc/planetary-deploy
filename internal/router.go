package internal

import (
	"net/http"

	"github.com/matthewzhaocc/planetary-deploy/internal/org"
	"github.com/matthewzhaocc/planetary-deploy/internal/types"
)

func GetMux(serv types.Server) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/org", org.AddOrgRoute(serv))
	return mux
}
