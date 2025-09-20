package user

import (
	"net/http"

	"github.com/shafi21064/ecom-go/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("POST /users", mngr.With(
		http.HandlerFunc(h.CreateUsers),
	))

	mux.Handle("POST /users/login", mngr.With(
		http.HandlerFunc(h.Login),
	))
}
