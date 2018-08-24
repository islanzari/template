package todos

import (
	"net/http"
	"strconv"

	"github.com/islanzari/template/internal/controller/middleware/reqctx"
	"github.com/islanzari/template/internal/request"
	"github.com/julienschmidt/httprouter"
)

// Create creates new user
func (h Handle) DeleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log := reqctx.Logger(r.Context())

	i, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.WithError(err).Error("Error during parsing id")
		request.Error(w, "Server error", http.StatusInternalServerError)
	}

	err = h.Todos.DeleteTodo(r.Context(), i)
	if err != nil {
		log.WithError(err).Error("Error during delete row")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	request.Success(w, nil)
	log.Info("User deleted todo")
}
