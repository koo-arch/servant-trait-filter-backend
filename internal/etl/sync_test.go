package etl_test

import (
	"context"
	"testing"
	"sort"

	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/internal/atlas"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/etl"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
)

func TestSyncServants(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
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
	ascRepo := repository.NewAscensionRepository(client)

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
		ascRepo,
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
	}
	// Verify that ascensions were synced correctly
	ascensions, err := ascRepo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(ascensions), 0, "Ascensions list should not be empty")
	for _, asc := range ascensions {
		assert.Greater(t, asc.ServantID, 0, "Ascension servant ID should be greater than 0")
		assert.Greater(t, asc.Stage, 0, "Ascension stage should be greater than 0")
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
	// Verify that the golden data matches the synced data
	for id, want := range golden {
        // Servant
        gotSvt, err := svtRepo.Get(ctx, id)
        assert.NoError(t, err, "servant %d should exist", id)

        // → Traits は順序を無視して比較
		entTraits := gotSvt.Edges.Traits
		var gotTraitIDs []int
		for _, t := range entTraits {
			gotTraitIDs = append(gotTraitIDs, t.ID)
		}
        sort.Ints(gotTraitIDs)
        sort.Ints(want.servant.Traits)
		
		gotSvtModel := model.Servant{
			ID:           gotSvt.ID,
			CollectionNo: gotSvt.CollectionNo,
			Name:         gotSvt.Name,
			Face:         gotSvt.Face,
			ClassID:      gotSvt.ClassID,
			Traits:       gotTraitIDs,
		}

        if diff := cmp.Diff(
            want.servant, gotSvtModel,

        ); diff != "" {
            t.Errorf("servant %d mismatch (-want +got):\n%s", id, diff)
        }

        // Ascension (1 再臨だけ登録している前提)
        gotAsc, err := ascRepo.GetByServantAndStage(ctx, id, 1)
        assert.NoError(t, err, "ascension s%d-st1 should exist", id)

		gotAscModel := model.Ascension{
			ServantID:   gotAsc.ServantID,
			Stage:       gotAsc.Stage,
			AttributeID: gotAsc.AttributeID,
			MoralAlignmentID: gotAsc.MoralAlignmentID,
			OrderAlignmentID: gotAsc.OrderAlignmentID,
		}

        if diff := cmp.Diff(
            want.ascension, gotAscModel,
        ); diff != "" {
            t.Errorf("ascension %d mismatch (-want +got):\n%s", id, diff)
        }
    }
}

var golden = map[int]struct {
    servant    model.Servant
    ascension  model.Ascension   // 1 再臨分しか取っていない前提
}{
    1001700: {
        servant: model.Servant{
            ID:           1001700,
            CollectionNo: 416,
            Name:         "ひびき＆千鍵",
            Face:         "https://static.atlasacademy.io/JP/Faces/f_10017000.png",
            ClassID:      10,
            Traits: []int{
                109, 405, 1000, 2000, 2001,
                2797, 2835, 2848, 5000,
            },
        },
        ascension: model.Ascension{
            ServantID:   1001700,
            Stage:       1,
            AttributeID: 202,
            // Alignment は無いので 0 / omitted
        },
    },
    800100: {
        servant: model.Servant{
            ID:           800100,
            CollectionNo: 1,
            Name:         "マシュ・キリエライト",
            Face:         "https://static.atlasacademy.io/JP/Faces/f_8001000.png",
            ClassID:      8,
            Traits: []int{
                2, 107, 1000,
                2001, 2009, 2631, 2654, 2780, 5000,
            },
        },
        ascension: model.Ascension{
            ServantID:        800100,
            Stage:            1,
            AttributeID:      201,
            OrderAlignmentID: 300,
            MoralAlignmentID: 303,
        },
    },
    505300: {
        servant: model.Servant{
            ID:           505300,
            CollectionNo: 385,
            Name:         "雨の魔女トネリコ",
            Face:         "https://static.atlasacademy.io/JP/Faces/f_5053000.png",
            ClassID:      5,
            Traits: []int{
                2, 104, 405, 1000,
                1177, 2001, 2007, 2008, 2011,
                2037, 2858, 2923, 5000,
            },
        },
        ascension: model.Ascension{
            ServantID:        505300,
            Stage:            1,
            AttributeID:      201,
            OrderAlignmentID: 300,
        },
    },
}