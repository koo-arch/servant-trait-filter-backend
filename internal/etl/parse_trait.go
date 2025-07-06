package etl

import (
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

type ParseTrait struct {
	attr, moral, order int
	others []int
}

func parseTraits(traits []model.Trait) ParseTrait {
	p := ParseTrait{
		attr: 0,
		moral: 0,
		order: 0,
		others: make([]int, 0),
	}
	for _, t := range traits {
		if t.ID == 0 || t.Name == UnknownTrait {
			continue
		}
		switch classifyTrait(t.ID) {
		case TraitAttribute:
			p.attr = t.ID
		case TraitOrderAlign:
			p.order = t.ID
		case TraitMoralAlign:
			p.moral = t.ID
		default:
			p.others = append(p.others, t.ID)
		}
	}


	return p
}