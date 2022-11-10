package tokenhandler

import "net/http"

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mock token response!\n"))
}
