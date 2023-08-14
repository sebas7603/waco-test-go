package models

type IndexResponse struct {
	Info    IndexInfo   `json:"info"`
	Results []Character `json:"results"`
}

type IndexInfo struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  any    `json:"prev"`
}
