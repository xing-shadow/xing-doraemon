package view

type PromList struct {
	Pagination `json:"pagination"`
	PromList   []PromItem `json:"list"`
}

type PromItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
