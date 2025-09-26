package product

import (
	"net/http"

	"e-com/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	product, err := h.svc.List()

	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, product, http.StatusOK)
}
