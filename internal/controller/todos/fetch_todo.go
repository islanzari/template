package todos

import (
	"net/http"
	"strconv"

	"github.com/islanzari/template/internal/controller/middleware/reqctx"
	"github.com/islanzari/template/internal/request"
	"github.com/julienschmidt/httprouter"
)

func (h Handle) FetchTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log := reqctx.Logger(r.Context())

	i, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.WithError(err).Error("Error bad convert string to int")
		request.Error(w, "Invalid string", http.StatusBadRequest)
		return
	}

	todo, err := h.Todos.FetchTodo(r.Context(), i)
	if err != nil {
		log.WithError(err).Error("Error while fetch todo")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	request.Success(w, todo)
}
