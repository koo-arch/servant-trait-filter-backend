package main

import (
	"context"
	"log"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/di"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/service"
)

func connectDB(ctx context.Context, url string) *ent.Client {
	client, err := ent.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func InitRepos(client *ent.Client) *di.Repos {
	// リポジトリの初期化
	svtRepo := repository.NewServantRepository(client)
	traitRepo := repository.NewTraitRepository(client)
	classRepo := repository.NewClassRepository(client)
	attrRepo := repository.NewAttributeRepository(client)
	orderAlignRepo := repository.NewOrderAlignmentRepository(client)
	moralAlignRepo := repository.NewMoralAlignmentRepository(client)
	ascensionRepo := repository.NewAscensionRepository(client)

	return &di.Repos{
		Servant: svtRepo,
		Trait: traitRepo,
		Class: classRepo,
		Attribute: attrRepo,
		OrderAlign: orderAlignRepo,
		MoralAlign: moralAlignRepo,
		Ascension: ascensionRepo,
	}
}

func InitServices(repos *di.Repos) *di.Services {
	// サービスの初期化
	svtService := service.NewServantServiceImpl(repos.Servant)
	traitService := service.NewTraitServiceImpl(repos.Trait)
	classService := service.NewClassServiceImpl(repos.Class)
	attrService := service.NewAttributeServiceImpl(repos.Attribute)
	orderAlignService := service.NewOrderAlignmentServiceImpl(repos.OrderAlign)
	moralAlignService := service.NewMoralAlignmentServiceImpl(repos.MoralAlign)

	return &di.Services{
		Servant: svtService,
		Trait: traitService,
		Class: classService,
		Attribute: attrService,
		OrderAlign: orderAlignService,
		MoralAlign: moralAlignService,
	}
}