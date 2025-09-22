package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shafi21064/ecom-go/repo"
	"github.com/shafi21064/ecom-go/util"
)

type RequestUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsOwner  bool   `json:"is_owner"`
}
func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {

	var newUser RequestUser

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userRepo.Create(repo.User{
		Name: newUser.Name,
		Email: newUser.Email,
		Password: newUser.Password,
		IsOwner: newUser.IsOwner,
	})

	util.SendData(w, createdUser, http.StatusCreated)
}
