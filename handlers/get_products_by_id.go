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

	for _, product := range database.ProductList {
		if productId == product.ID {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}

	util.SendData(w, "Product not found", http.StatusNotFound)

}
