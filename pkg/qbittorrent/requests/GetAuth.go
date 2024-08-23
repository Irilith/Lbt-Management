package requests

import (
	v2 "Lbt-Management/pkg/qbittorrent/api/v2"
	"fmt"
	"net/http"
)

func GetAuth(username, password string) (string, error) {
	client := v2.NewClient("auth/login")
	params := map[string]string{
		"username": username,
		"password": password,
	}
	req, err := client.NewRequest(http.MethodPost, params)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	var sid string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "SID" {
			sid = cookie.Value
			break
		}
	}

	return sid, nil
}
