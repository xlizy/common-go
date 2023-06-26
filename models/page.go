package models

type Page struct {
	PageNum  int `json:"pageNum" url:"pageNum" gorm:"-"`
	PageSize int `json:"pageSize" url:"pageSize" gorm:"-"`
}

func (p Page) GetLimit() int {
	return p.PageSize
}

func (p Page) GetOffset() int {
	return (p.PageNum - 1) * p.PageSize
}
