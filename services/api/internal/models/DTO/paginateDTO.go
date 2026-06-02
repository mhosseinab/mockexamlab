package DTO

type PaginateResponse struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
