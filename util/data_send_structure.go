package util

import (
	"net/http"
)

type CommonStructure struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type PaginationStructure struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalItem int `json:"total_item"`
	TotalPage int `json:"total_page"`
}

type PaginatedStructure struct {
	Result     bool                `json:"result"`
	Message    string              `json:"message"`
	Data       any                 `json:"data"`
	Pagination PaginationStructure `json:"pagination"`
}

func SendCommonData(w http.ResponseWriter, result bool, message string, data any, statusCode int) {
	SendData(w, CommonStructure{
		Result:  result,
		Message: message,
		Data:    data,
	}, statusCode)
}

func SendPaginatedDataData(w http.ResponseWriter, data any, page, limit, totalItem, totalPage int) {
	SendData(w, PaginatedStructure{
		Result:  true,
		Message: "Data Fetched successfully",
		Data:    data,
		Pagination: PaginationStructure{
			Page:      page,
			Limit:     limit,
			TotalItem: totalItem,
			TotalPage: totalPage,
		},
	}, http.StatusOK)
}
