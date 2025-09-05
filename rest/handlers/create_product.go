package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shafi21064/ecom-go/database"
	"github.com/shafi21064/ecom-go/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, http.StatusCreated)
}
