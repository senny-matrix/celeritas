package handlers

import (
	"github.com/senny-matrix/celeritas"
	"github.com/senny-matrix/coolapp/data"
	"net/http"
)

type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}
