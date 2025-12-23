package types

type SearchModsResponse struct {
	Data       []Data     `json:"data"`
	Pagination Pagination `json:"pagination"`
}
