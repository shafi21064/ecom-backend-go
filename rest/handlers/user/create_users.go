package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shafi21064/ecom-go/database"
	"github.com/shafi21064/ecom-go/util"
)

func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {

	var newUser database.Users

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	util.SendData(w, createdUser, http.StatusCreated)
}
