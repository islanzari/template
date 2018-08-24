package todos

import (
	"encoding/json"
	"net/http"

	"github.com/islanzari/template/internal/controller/middleware/reqctx"
	"github.com/islanzari/template/internal/request"
	"github.com/julienschmidt/httprouter"
)

func (h Handle) CreateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log := reqctx.Logger(r.Context())

	var Data struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		log.WithError(err).Error("Error while decoding request body")
		request.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	todo, err := h.Todos.CreateTodo(r.Context(), Data.Name, Data.Description)
	if err != nil {
		log.WithError(err).Error("Error while creating todo")
		request.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	request.Success(w, todo)
	log.Info("user created todo")
}
