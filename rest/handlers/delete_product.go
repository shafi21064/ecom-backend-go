package handlers

import (
	"net/http"
	"strconv"

	"github.com/shafi21064/ecom-go/database"
	"github.com/shafi21064/ecom-go/util"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productIDString := r.PathValue("id")

	productId, err := strconv.Atoi(productIDString)

	if err != nil {
		http.Error(w, "Please enter a valid id", http.StatusBadRequest)
		return
	}

	database.Delete(productId)

	util.SendData(w, "Successfully deleted", http.StatusOK)

}
