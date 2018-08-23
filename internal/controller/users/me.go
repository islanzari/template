package users

import (
	"net/http"

	"github.com/islanzari/template/internal/controller/middleware/reqctx"
	"github.com/islanzari/template/internal/request"
	"github.com/julienschmidt/httprouter"
)

// Me returns info of user's by id provided in authorization token
func (h Handle) Me(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params) {
	log := reqctx.Logger(r.Context())
	user := reqctx.User(r.Context())

	log.Info("User fetched his information")
	request.Success(w, user)
}
