package di

import (
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/service"
)

type Repos struct {
	Servant repository.ServantRepository
	Trait   repository.TraitRepository
	Class   repository.ClassRepository
	Attribute repository.AttributeRepository
	OrderAlign repository.OrderAlignmentRepository
	MoralAlign repository.MoralAlignmentRepository
	Ascension repository.AscensionRepository
}

type Services struct {
	Servant service.ServantService
	Trait   service.TraitService
	Class   service.ClassService
	Attribute service.AttributeService
	OrderAlign service.OrderAlignmentService
	MoralAlign service.MoralAlignmentService
}