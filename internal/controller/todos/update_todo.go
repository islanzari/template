package todos

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/islanzari/template/internal/controller/middleware/reqctx"
	"github.com/islanzari/template/internal/request"
	"github.com/julienschmidt/httprouter"
)

func (h Handle) UpdateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log := reqctx.Logger(r.Context())

	var data struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.WithError(err).Error("Error while decoding request body")
		request.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	i, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.WithError(err).Error("Error bad convert string to int")
		request.Error(w, "Invalid string", http.StatusBadRequest)
		return
	}

	err = h.Todos.UpdateTodo(r.Context(), i, data.Name, data.Description)
	if err != nil {
		log.WithError(err).Error("Error while updateing todo")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

}
