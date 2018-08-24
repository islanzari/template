package todos

import (
	"net/http"

	"github.com/islanzari/template/internal/controller/middleware/reqctx"
	"github.com/islanzari/template/internal/request"
	"github.com/julienschmidt/httprouter"
)

func (h Handle) FetchTodos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log := reqctx.Logger(r.Context())

	todo, err := h.Todos.FetchTodos(r.Context())
	if err != nil {
		log.WithError(err).Error("Error while fetch todos")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	request.Success(w, todo)
}
