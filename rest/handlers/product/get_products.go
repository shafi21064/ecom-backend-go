package product

import (
	"log"
	"net/http"
	"strconv"

	"e-com/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	quaryParam := r.URL.Query()

	pageString := quaryParam.Get("page")

	limitString := quaryParam.Get("limit")

	pageNumber, _ := strconv.Atoi(pageString)
	limitNumber, _ := strconv.Atoi(limitString)

	if pageNumber == 0 {
		pageNumber = 1
	}
	if limitNumber == 0 {
		limitNumber = 10
	}
	product, err := h.svc.List(pageNumber, limitNumber)

	if err != nil {
		log.Print(err.Error())
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	totalCount, err := h.svc.Count()
	if err != nil {
		log.Print(err.Error())
		util.SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedDataData(
		w,
		product,
		pageNumber,
		limitNumber,
		totalCount,
		totalCount/limitNumber,
	)
}
