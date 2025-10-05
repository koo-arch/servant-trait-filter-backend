package search

type ServantSearchQuery struct {
	Root Expr `json:"root"`
	Limit int  `json:"limit"`
	Offset int  `json:"offset"`
}