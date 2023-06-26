package models

type SelectOptions struct {
	Value int    `json:"value"`
	Desc  string `json:"desc"`
}

type CharSelectOptions struct {
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type ViewEnumCache struct {
	Val  int    `json:"val"`
	Des  string `json:"des"`
	Enum string `json:"enum"`
}
