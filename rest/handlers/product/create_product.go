package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shafi21064/ecom-go/repo"
	"github.com/shafi21064/ecom-go/util"
)

type RequestProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct RequestProduct

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error", err)
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	})

	if err != nil {

		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}
