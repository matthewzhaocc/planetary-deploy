package internal

import "net/http"

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
