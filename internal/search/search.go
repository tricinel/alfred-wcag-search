// Package search provides the search functionality
package search

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/sahilm/fuzzy"
	"github.com/tricinel/alfred-wcag-search/internal/models"
)

func Find(query string, data []byte) ([]models.WCAGItem, error) {
	var items []models.WCAGItem
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	if query == "" {
		return items, nil
	}

	var finalResults []models.WCAGItem
	seenIDs := make(map[string]bool)

	// Standard Fuzzy Search
	targets := make([]string, len(items))
	for i, item := range items {
		targets[i] = fmt.Sprintf("%s %s %s", item.Title, item.Keywords, item.Level)
	}

	matches := fuzzy.Find(query, targets)
	for _, match := range matches {
		item := items[match.Index]
		finalResults = append(finalResults, item)
		seenIDs[item.ID] = true
	}

	// Levenshtein (Only if item wasn't already found)
	query = strings.ToLower(query)
	for _, item := range items {
		if seenIDs[item.ID] {
			continue
		}

		words := strings.Fields(strings.ToLower(item.Title))
		words = append(words, item.Keywords)

		for _, word := range words {
			if levenshtein.ComputeDistance(query, word) <= 2 {
				// Mark as a suggestion
				item.Title = "Did you mean: " + item.Title
				finalResults = append(finalResults, item)
				break
			}
		}
	}

	return finalResults, nil
}
