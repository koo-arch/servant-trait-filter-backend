package etl_test

import (
	"context"
	"testing"

	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/internal/atlas"
	"github.com/koo-arch/servant-trait-filter-backend/internal/etl"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
)

func TestSyncServants(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	client = client.Debug()
	defer client.Close()

	err := client.Schema.Create(ctx)
	require.NoError(t, err)
	
	// Create repositories
	classRepo := repository.NewClassRepository(client)
	attrRepo := repository.NewAttributeRepository(client)
	moralRepo := repository.NewMoralAlignmentRepository(client)
	orderRepo := repository.NewOrderAlignmentRepository(client)
	svtRepo := repository.NewServantRepository(client)
	traitRepo := repository.NewTraitRepository(client)

	// Create a mock Atlas client
	atlasClient := atlas.NewClient()
	syncer := etl.NewSyncAtlas(
		client,
		atlasClient,
		classRepo,
		attrRepo,
		moralRepo,
		orderRepo,
		svtRepo,
		traitRepo,
	)
	err = syncer.Sync(ctx)
	require.NoError(t, err)
	

	// Verify that servants were synced correctly
	servants, err := svtRepo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(servants), 0, "Servants list should not be empty")
	for _, svt := range servants {
		assert.NotEmpty(t, svt.Name, "Servant name should not be empty")
		assert.NotEmpty(t, svt.Face, "Servant face should not be empty")
		assert.Greater(t, svt.CollectionNo, 0, "Servant collection number should be greater than 0")
		assert.Greater(t, svt.ClassID, 0, "Servant class ID should be greater than 0")
		assert.Greater(t, svt.AttributeID, 0, "Servant attribute ID should be greater than 0")
	}
	// Verify that traits were synced correctly
	traits, err := traitRepo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(traits), 0, "Traits list should not be empty")
	for _, trait := range traits {
		assert.NotEmpty(t, trait.NameEn, "Trait name should not be empty")
		assert.Greater(t, trait.ID, 0, "Trait ID should be greater than 0")
	}
	// Verify that classes were synced correctly
	classes, err := classRepo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(classes), 0, "Classes list should not be empty")
	for _, class := range classes {
		assert.NotEmpty(t, class.NameEn, "Class name should not be empty")
		assert.Greater(t, class.ID, 0, "Class ID should be greater than 0")
	}
	// Verify that attributes were synced correctly
	attributes, err := attrRepo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(attributes), 0, "Attributes list should not be empty")
	for _, attr := range attributes {
		assert.NotEmpty(t, attr.NameEn, "Attribute name should not be empty")
		assert.Greater(t, attr.ID, 0, "Attribute ID should be greater than 0")
	}
	// Verify that order alignments were synced correctly
	orderAlignments, err := orderRepo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(orderAlignments), 0, "Order alignments list should not be empty")
	for _, order := range orderAlignments {
		assert.NotEmpty(t, order.NameEn, "Order alignment name should not be empty")
		assert.Greater(t, order.ID, 0, "Order alignment ID should be greater than 0")
	}
	// Verify that moral alignments were synced correctly
	moralAlignments, err := moralRepo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(moralAlignments), 0, "Moral alignments list should not be empty")
	for _, moral := range moralAlignments {
		assert.NotEmpty(t, moral.NameEn, "Moral alignment name should not be empty")
		assert.Greater(t, moral.ID, 0, "Moral alignment ID should be greater than 0")
	}
}