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

type ClassDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AttributeDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MoralAlignmentDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OrderAlignmentDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TraitDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}