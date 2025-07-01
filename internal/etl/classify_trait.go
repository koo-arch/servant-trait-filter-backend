package etl

type TraitKind int

const (
	TraitOther TraitKind = iota
	TraitAttribute
	TraitOrderAlign
	TraitMoralAlign
)

func classifyTrait(traitID int) TraitKind {
	switch {
	case traitID >= 200 && traitID < 300:
		return TraitAttribute
	case traitID >= 300 && traitID < 303:
		return TraitOrderAlign
	case traitID >= 303 && traitID < 400:
		return TraitMoralAlign
	default:
		return TraitOther
	}
}