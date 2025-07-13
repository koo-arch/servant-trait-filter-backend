package search

type Expr struct {
    And  []*Expr `json:"and,omitempty"`
    Or   []*Expr `json:"or,omitempty"`
    Not  *Expr   `json:"not,omitempty"`

    // 原子条件
    TraitID         *int     `json:"trait,omitempty"`
    ClassID         *int     `json:"class,omitempty"`
    AttributeID     *int     `json:"attribute,omitempty"`
    OrderAlignID    *int     `json:"orderAlignment,omitempty"`
    MoralAlignID    *int     `json:"moralAlignment,omitempty"`
}