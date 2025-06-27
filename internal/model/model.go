package model

type Servant struct {
	ID               int     `json:"id"`
	CollectionNo     int     `json:"collectionNo"`
	Name             string  `json:"name"`
	Face             string  `json:"face"`
	ClassID          int     `json:"classId"`
	AttributeID      int     `json:"attributeId"`
	MoralAlignmentID int     `json:"moralAlignmentId,omitempty"`
	OrderAlignmentID int     `json:"orderAlignmentId,omitempty"`
	Traits           []int   `json:"traits"`
}

type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Attribute struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OrderAlignment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MoralAlignment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}