package product

import (
	"net/http"
	"strconv"

	"e-com/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productIDString := r.PathValue("id")

	productId, err := strconv.Atoi(productIDString)

	if err != nil {
		util.SendError(w, "Please enter a valid id", http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(productId)

	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Successfully deleted", http.StatusOK)

}
