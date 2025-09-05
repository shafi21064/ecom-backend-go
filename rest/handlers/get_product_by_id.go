package handlers

import (
	"net/http"
	"strconv"

	"github.com/shafi21064/ecom-go/database"
	"github.com/shafi21064/ecom-go/util"
)

func GetProductsByID(w http.ResponseWriter, r *http.Request) {

	productIDString := r.PathValue("id")

	productId, err := strconv.Atoi(productIDString)

	if err != nil {
		http.Error(w, "Please enter a valid id", http.StatusBadRequest)
		return
	}

	product := database.Get(productId)

	if product == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}
	util.SendData(w, product, http.StatusOK)
}
