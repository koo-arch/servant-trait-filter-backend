package search

type ServantSearchQuery struct {
	Root Expr
	Limit int
	Offset int
}