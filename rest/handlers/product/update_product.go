package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"e-com/repo"
	"e-com/util"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct RequestProduct

	productIDString := r.PathValue("id")

	productId, err := strconv.Atoi(productIDString)

	if err != nil {
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
	}

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	updatedProduct, err := h.productRepo.Update(
		productId,
		repo.Product{
			Title:       newProduct.Title,
			Description: newProduct.Description,
			Price:       newProduct.Price,
			ImgUrl:      newProduct.ImgUrl,
		})

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if updatedProduct == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, updatedProduct, http.StatusOK)
}
