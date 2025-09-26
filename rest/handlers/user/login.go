package user

import (
	"encoding/json"
	"net/http"

	"e-com/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&reqLogin)

	if err != nil {
		println(err)
		util.SendError(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr, err := h.svc.Get(reqLogin.Email, reqLogin.Password)

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if usr == nil {
		util.SendError(w, "Invalid Credential", http.StatusBadRequest)
		return
	}

	jwt, err := util.CreateJwt(
		h.cnf.JwtSecrateKey,
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
