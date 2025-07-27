package model

import (
	. "github.com/koo-arch/servant-trait-filter-backend/internal/search"
)

type ServantDTO struct {
	ID               int      `json:"id"`
	CollectionNo     int      `json:"collectionNo"`
	Name             string   `json:"name"`
	Face             string   `json:"face"`
}

type SearchResponseDTO struct {
	Total   int          `json:"total"`
	Offset  int          `json:"offset"`
	Limit   int          `json:"limit"`
	Items   []ServantDTO `json:"items"`
}

type SearchRequestDTO struct {
	Root   Expr   `json:"root" binding:"required"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
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