// Unused.
// I made this when i giving up making regexp, but i found solution later on
package anilist

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type AniListResponse struct {
	Data struct {
		Media struct {
			Title struct {
				Romaji  string `json:"romaji"`
				English string `json:"english"`
				Native  string `json:"native"`
			} `json:"title"`
		} `json:"Media"`
	} `json:"data"`
}

func SearchAniList(query string) (string, error) {
	client := resty.New()
	queryString := `query ($search: String) {
		Media (search: $search, type: MANGA) {
			title {
				romaji
				english
				native
			}
		}
	}`

	response, err := client.R().
		SetBody(map[string]interface{}{
			"query": queryString,
			"variables": map[string]string{
				"search": query,
			},
		}).
		SetHeader("Content-Type", "application/json").
		Post("https://graphql.anilist.co")
	if err != nil {
		return "", err
	}

	var result AniListResponse
	if err := json.Unmarshal(response.Body(), &result); err != nil {
		return "", err
	}

	// Return the English title if available, otherwise fall back to romaji or native title
	if result.Data.Media.Title.English != "" {
		return result.Data.Media.Title.English, nil
	} else if result.Data.Media.Title.Romaji != "" {
		return result.Data.Media.Title.Romaji, nil
	} else if result.Data.Media.Title.Native != "" {
		return result.Data.Media.Title.Native, nil
	}

	return "No match found", nil
}
