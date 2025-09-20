package product

import (
	"net/http"

	"github.com/shafi21064/ecom-go/database"
	"github.com/shafi21064/ecom-go/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), http.StatusOK)
}
