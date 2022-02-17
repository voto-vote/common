package common

type Links struct {
	Self string `json:"self"`
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type MetaData struct {
	Count int   `json:"count"`
	Total int   `json:"total"`
	Page  int   `json:"page"`
	Links Links `json:"_links"`
}
