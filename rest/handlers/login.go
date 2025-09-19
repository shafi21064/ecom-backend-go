package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shafi21064/ecom-go/config"
	"github.com/shafi21064/ecom-go/database"
	"github.com/shafi21064/ecom-go/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&reqLogin)

	if err != nil {
		println(err)
		util.SendError(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)

	if usr == nil {
		util.SendError(w, "Invalid Credential", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()

	jwt, err := util.CreateJwt(
		cnf.JwtSecrateKey,
		util.Payload{
			Sub:         usr.ID,
			Name:        usr.Name,
			Email:       usr.Email,
			IsShopOwner: usr.IsOwner,
		},
	)

	if err != nil {
		print(err)
		return
	}

	util.SendData(w, jwt, http.StatusOK)

}
