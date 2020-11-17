package view

type PaginationRequest struct {
	Page     uint `query:"page"`
	PageSize uint `query:"page_size"`
}

type PaginationResp struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
}
