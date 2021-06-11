package view

type Pagination struct {
	Total       int `json:"total"`
	PageSize    int `query:"page_size"`
	CurrentPage int `json:"current_page"`
}
