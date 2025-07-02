package atlas

import (
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

type Servant struct {
	ID          int     `json:"id"`
	CollectionNo int     `json:"collectionNo"`
	Type		string  `json:"type"`
	Name        string  `json:"name"`
	Face        string  `json:"face"`
	ClassID     int     `json:"classId"`
	Class       string  `json:"className"`
	Attribute   string  `json:"attribute"`
	Traits	    []model.Trait `json:"traits"`
}
