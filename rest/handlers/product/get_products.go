package product

import (
	"net/http"

	"github.com/shafi21064/ecom-go/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	product, err := h.productRepo.List()

	if err != nil{
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, product, http.StatusOK)
}
