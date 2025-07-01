package atlas

type Servant struct {
	ID          int     `json:"id"`
	CollectionNo int     `json:"collectionNo"`
	Name        string  `json:"name"`
	Face        string  `json:"face"`
	ClassID     int     `json:"classId"`
	Class       string  `json:"className"`
	Attribute   string  `json:"attribute"`
	Traits	    []Trait `json:"traits"`
}

type Trait struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}