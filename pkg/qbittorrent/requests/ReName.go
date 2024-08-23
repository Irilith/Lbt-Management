package requests

import (
	v2 "Lbt-Management/pkg/qbittorrent/api/v2"
	"fmt"
	"net/http"
)

func RenameFolder(sid string, hash string, oldPath string, newPath string) error {
	client := v2.NewClient("torrents/renameFolder")
	params := map[string]string{
		"hash":    hash,
		"oldPath": oldPath,
		"newPath": newPath,
	}
	req, err := client.NewRequest(http.MethodPost, params)
	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{
		Name:  "SID",
		Value: sid,
	})
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

func RenameFile(sid string, hash string, oldPath string, newPath string) error {
	client := v2.NewClient("torrents/renameFile")
	params := map[string]string{
		"hash":    hash,
		"oldPath": oldPath,
		"newPath": newPath,
	}
	req, err := client.NewRequest(http.MethodPost, params)
	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{
		Name:  "SID",
		Value: sid,
	})
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
