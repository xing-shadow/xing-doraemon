package view

type PaginationRequest struct {
	Page     uint `form:"page" json:"page" bind:"required"`
	PageSize uint `form:"page_size" json:"page_size"`
}

type PaginationResp struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
}
