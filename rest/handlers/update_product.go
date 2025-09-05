package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shafi21064/ecom-go/database"
	"github.com/shafi21064/ecom-go/util"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	updatedProduct := database.Update(newProduct)
	if updatedProduct == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, updatedProduct, http.StatusOK)
}
