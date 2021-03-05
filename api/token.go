package api

import (
	"fmt"
	"net/http"

	"github.com/Xuanwo/utterances-oauth.go/cors"
)

func Token(w http.ResponseWriter, r *http.Request) {
	cors.Add(w, r)
	_cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, _ = fmt.Fprintf(w, `"%s"`, _cookie.Value)
}
