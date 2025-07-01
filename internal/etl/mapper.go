package etl

import (
	"sort"

	"github.com/koo-arch/servant-trait-filter-backend/internal/atlas"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

const (
	EnemyCollectionDetail = "enemyCollectionDetail"
	UnknownTrait          = "unknown"
)

func (s *SyncAtlas) extractClass(atlasData []atlas.Servant) []model.Class {
	classMap := make(map[int]string)
	for _, servant := range atlasData {
		classMap[servant.ClassID] = servant.Class
	}
	classes := make([]model.Class, 0, len(classMap))
	for id, name := range classMap {
		classes = append(classes, model.Class{
			ID:   id,
			Name: name,
		})
	}
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].ID < classes[j].ID
	})
	return classes
}

func (s *SyncAtlas) extractPlayableServants(atlasData []atlas.Servant) []model.Servant {
	servants := make([]model.Servant, 0)
	for _, servant := range atlasData {
		if servant.Type != EnemyCollectionDetail {
			servants = append(servants, s.toModelServant(servant))
		}
	}

	sort.Slice(servants, func(i, j int) bool {
		return servants[i].ID < servants[j].ID
	})
	return servants
}

func (s *SyncAtlas) toModelServant(servant atlas.Servant) model.Servant {
	var orderID, moralID, attrID int
	traitIDs := make([]int, 0)
	for _, trait := range servant.Traits {
		if trait.ID == 0 || trait.Name == UnknownTrait {
			continue
		}
		switch classifyTrait(trait.ID) {
		case TraitAttribute:
			attrID = trait.ID
		case TraitOrderAlign:
			orderID = trait.ID
		case TraitMoralAlign:
			moralID = trait.ID
		default:
			traitIDs = append(traitIDs, trait.ID)
		}
	}

	return model.Servant{
		ID:               servant.ID,
		CollectionNo:     servant.CollectionNo,
		Name:             servant.Name,
		Face:             servant.Face,
		ClassID:          servant.ClassID,
		OrderAlignmentID: orderID,
		MoralAlignmentID: moralID,
		AttributeID:      attrID,
		Traits:           traitIDs,
	}
}

func (s *SyncAtlas) extractMetaFromTraits(traits []model.Trait) ([]model.Attribute, []model.OrderAlignment, []model.MoralAlignment) {
	attributes := make([]model.Attribute, 0)
	orderAlignments := make([]model.OrderAlignment, 0)
	moralAlignments := make([]model.MoralAlignment, 0)

	for _, trait := range traits {
		switch classifyTrait(trait.ID) {
		case TraitAttribute:
			attributes = append(attributes, model.Attribute{ID: trait.ID, Name: trait.Name})
		case TraitOrderAlign:
			orderAlignments = append(orderAlignments, model.OrderAlignment{ID: trait.ID, Name: trait.Name})
		case TraitMoralAlign:
			moralAlignments = append(moralAlignments, model.MoralAlignment{ID: trait.ID, Name: trait.Name})
		}
	}

	sort.Slice(attributes, func(i, j int) bool { return attributes[i].ID < attributes[j].ID })
	sort.Slice(orderAlignments, func(i, j int) bool { return orderAlignments[i].ID < orderAlignments[j].ID })
	sort.Slice(moralAlignments, func(i, j int) bool { return moralAlignments[i].ID < moralAlignments[j].ID })

	return attributes, orderAlignments, moralAlignments
}