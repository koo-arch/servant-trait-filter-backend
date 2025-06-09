package atlas

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

func (c *client) FetchServants(ctx context.Context, region string) ([]Servant, error) {
	rel := &url.URL{
		Path: fmt.Sprintf("/export/%s/basic_servant.json", region),
	}
	var list []Servant
	if err := c.DoJSON(ctx, http.MethodGet, rel, nil, &list); err != nil {
		return nil, fmt.Errorf("failed to fetch servant: %w", err)
	}
	
	return list, nil
}

func (c *client) FetchTraits(ctx context.Context, region string) ([]Trait, error) {
	rel := &url.URL{
		Path: fmt.Sprintf("/export/%s/nice_trait.json", region),
	}
	var traits map[string]string
	if err := c.DoJSON(ctx, http.MethodGet, rel, nil, &traits); err != nil {
		return nil, fmt.Errorf("failed to fetch traits: %w", err)
	}

	traitList, err := toTraitSlice(traits)
	if err != nil {
		return nil, fmt.Errorf("failed to convert traits: %w", err)
	}
	return traitList, nil
}

func toTraitSlice(traits map[string]string) ([]Trait, error) {
	traitList := make([]Trait, 0, len(traits))
	for id, name := range traits {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("invalid trait ID %s: %w", id, err)
		}
		traitList = append(traitList, Trait{
			ID:   idInt,
			Name: name,
		})
	}
	sort.Slice(traitList, func(i, j int) bool {
		return traitList[i].ID < traitList[j].ID
	})

	return traitList, nil
}