package view

type GetProms struct {
	Page     uint   `json:"page"`
	PageSize uint   `json:"page_size"`
	Name     string `query:"name"`
	Url      string `query:"url"`
}

type GetProm struct {
	ID uint8 `query:"Id"`
}

type CreateProm struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required"`
}

type ModifyProm struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type DeleteProm struct {
	ID uint `json:"id" binding:"required"`
}
