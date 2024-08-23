package requests

import (
	v2 "Lbt-Management/pkg/qbittorrent/api/v2"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Files struct {
	// Availability int    `json:"availability"`
	// Index        int    `json:"index"`
	// IsSeed       bool   `json:"is_seed,omitempty"`
	Name string `json:"name"`
	// PieceRange   [2]int `json:"piece_range"`
	// Priority     int    `json:"priority"`
	// Progress     int    `json:"progress"`
	// Size         int    `json:"size"`
}

func GetFiles(sid string, hash string) []Files {
	client := v2.NewClient("torrents/files")
	params := map[string]string{
		"hash": hash,
	}
	req, err := client.NewRequest(http.MethodGet, params)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "SID",
		Value: sid,
	})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var files []Files
	if err := json.Unmarshal(body, &files); err != nil {
		panic(err)
	}

	return files
}
