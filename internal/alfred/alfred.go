// Package alfred provides the functions for the alfred workflow
package alfred

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tricinel/alfred-wcag-search/internal/models"
)

func Error(err error) {
	resp := models.AlfredResponse{
		Items: []models.AlfredItem{
			{
				Title:    "Error searching WCAG",
				Subtitle: err.Error(),
				Valid:    false,
				Icon:     &models.AlfredIcon{Path: "error.png"},
			},
		},
	}
	out, _ := json.Marshal(resp)
	fmt.Println(string(out))
}

func Empty() {
	resp := models.AlfredResponse{
		Items: []models.AlfredItem{
			{
				Title:    "Search the web accessibility guidelines (WCAG 2.2)",
				Subtitle: "e.g. keyboard. When the search is empty, you can open WCAG in your browser",
				Arg:      "https://www.w3.org/WAI/WCAG22/Understanding",
				Valid:    true,
			},
		},
	}
	out, _ := json.Marshal(resp)
	fmt.Println(string(out))
}

func NoResults(query string) {
	resp := models.AlfredResponse{
		Items: []models.AlfredItem{
			{
				Title:    fmt.Sprintf("Oops! Couldn't find anything for %s", query),
				Subtitle: "Maybe refine your search? You can search for keywords like \"keyboard\" or \"label\"",
				Valid:    false,
				Icon:     &models.AlfredIcon{Path: "warn.png"},
			},
		},
	}
	out, _ := json.Marshal(resp)
	fmt.Println(string(out))
}

func Format(items []models.WCAGItem) models.AlfredResponse {
	resp := models.AlfredResponse{Items: []models.AlfredItem{}}

	for _, res := range items {
		url := "https://www.w3.org/WAI/WCAG22/Understanding/" + res.Slug
		iconPath := "icon.png"

		if strings.HasPrefix(res.Title, "Did you mean") {
			iconPath = "info.png"
		}

		resp.Items = append(resp.Items, models.AlfredItem{
			Title:    res.Title,
			Subtitle: fmt.Sprintf("Level %s: %s", res.Level, url),
			Arg:      url,
			Match:    fmt.Sprintf("%s %s %s", res.Title, res.Keywords, res.Level),
			Valid:    true,
			Icon:     &models.AlfredIcon{Path: iconPath},
		})
	}

	return resp
}
