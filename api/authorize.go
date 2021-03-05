package api

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/Xuanwo/utterances-oauth.go/oauth"
	"github.com/Xuanwo/utterances-oauth.go/state"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	var query struct {
		RedirectUri string `form:"redirect_uri" binding:"required"`
	}
	if err := binding.Query.Bind(r, &query); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_state, err := state.EncryptState(query.RedirectUri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, oauth.AuthURL(_state), http.StatusFound)
}
