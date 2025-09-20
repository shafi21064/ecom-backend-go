package review

import (
	"net/http"

	"github.com/shafi21064/ecom-go/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, mngr middleware.Manager) {

	mux.Handle(
		"POST /reviews", mngr.With(http.HandlerFunc(h.GetReviews)))

}
