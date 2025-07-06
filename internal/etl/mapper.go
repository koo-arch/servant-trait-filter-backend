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

func (s *SyncAtlas) extractPlayable(atlasData []atlas.Servant) ([]model.Servant, []model.Ascension) {
	servants := make([]model.Servant, 0)
	ascensions := make([]model.Ascension, 0)
	for _, servant := range atlasData {
		if servant.Type != EnemyCollectionDetail {
			svt, asc := s.buildServantAndAsc(servant)
			servants = append(servants, svt)
			ascensions = append(ascensions, asc)
		}
	}

	sort.Slice(servants, func(i, j int) bool {
		return servants[i].ID < servants[j].ID
	})
	sort.Slice(ascensions, func(i, j int) bool {
		return ascensions[i].ServantID < ascensions[j].ServantID
	})
	return servants, ascensions
}

func (s *SyncAtlas) buildServantAndAsc(servant atlas.Servant) (model.Servant, model.Ascension) {
	p := parseTraits(servant.Traits)

	svt := model.Servant{
		ID:           servant.ID,
		CollectionNo: servant.CollectionNo,
		Name:         servant.Name,
		Face:         servant.Face,
		ClassID:      servant.ClassID,
		Traits:       p.others,
	}

	asc := model.Ascension{
		ServantID: 		  servant.ID,
		Stage:	          1,
		AttributeID:      p.attr,
		MoralAlignmentID: p.moral,
		OrderAlignmentID: p.order,
	}
	return svt, asc
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