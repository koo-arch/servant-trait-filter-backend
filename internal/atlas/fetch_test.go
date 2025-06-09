package atlas_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/koo-arch/servant-trait-filter-backend/internal/atlas"
	"github.com/stretchr/testify/assert"
)

func TestFetchServants(t *testing.T) {
	cli := atlas.NewClient()
	svts, err := cli.FetchServants(context.Background(), "JP")
	assert.NoError(t, err)
	assert.NotEmpty(t, svts)
	assert.Greater(t, len(svts), 0, "Servants list should not be empty")

	t.Log("Servants:", prettyJSON(svts[:10]))
}

func TestFetchTraits(t *testing.T) {
	cli := atlas.NewClient()
	traits, err := cli.FetchTraits(context.Background(), "JP")
	assert.NoError(t, err)
	assert.NotEmpty(t, traits)
	assert.Greater(t, len(traits), 0, "Traits list should not be empty")

	// Check if traits are sorted by ID
	for i := 1; i < len(traits); i++ {
		assert.LessOrEqual(t, traits[i-1].ID, traits[i].ID, "Traits should be sorted by ID")
	}

	t.Log("Traits:", prettyJSON(traits[:10]))
}

func prettyJSON(v any) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "Error marshalling to JSON"
	}
	return string(data)
}