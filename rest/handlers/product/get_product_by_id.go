package product

import (
	"net/http"
	"strconv"

	"e-com/util"
)

func (h *Handler) GetProductsByID(w http.ResponseWriter, r *http.Request) {

	productIDString := r.PathValue("id")

	productId, err := strconv.Atoi(productIDString)

	if err != nil {
		util.SendError(w, "Please enter a valid id", http.StatusBadRequest)
		return
	}

	product, err := h.svc.Get(productId)

	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if product == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}
	util.SendData(w, product, http.StatusOK)
}
