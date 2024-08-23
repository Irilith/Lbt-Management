package requests

import (
	v2 "Lbt-Management/pkg/qbittorrent/api/v2"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TorrentInfo struct {
	Hash        string `json:"hash"`
	SavePath    string `json:"save_path"`
	ContentPath string `json:"content_path"`
	Name        string `json:"name"`
}

func GetTorrentsInfos(sid string) ([]TorrentInfo, error) {
	client := v2.NewClient("torrents/info")
	req, err := client.NewRequest(http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "SID",
		Value: sid,
	})

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var torrents []TorrentInfo
	if err := json.Unmarshal(body, &torrents); err != nil {
		return nil, err
	}

	return torrents, nil
}
