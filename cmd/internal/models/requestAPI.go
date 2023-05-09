package models

type RequestAPI struct {
	CurrentPage int    `json:"current_Page"`
	Data        []Mod  `json:"data"`
	PrevPage    string `json:"prev_page"`
	NextPage    string `json:"next_page"`
	TotalPage   int    `json:"total_Page"`
}

func NewRequestAPI(currentPage int, data []Mod, prevPage string, nextPage string, totalPage int) *RequestAPI {
	return &RequestAPI{CurrentPage: currentPage, Data: data, PrevPage: prevPage, NextPage: nextPage, TotalPage: totalPage}
}
