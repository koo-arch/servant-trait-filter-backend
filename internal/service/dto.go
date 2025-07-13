package service

type ServantDTO struct {
	ID               int      `json:"id"`
	CollectionNo     int      `json:"collectionNo"`
	Name             string   `json:"name"`
	Face             string   `json:"face"`
}

type SearchResponseDTO struct {
	Total    int          `json:"total"`
	Offset  int          `json:"offset"`
	Limit   int          `json:"limit"`
	Items []ServantDTO `json:"items"`
}