package integrations

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// now i wil use chat gpt here because i need to take a nap
type UserStats struct {
	Data struct {
		Languages []struct {
			Name string `json:"name"`
			Time string `json:"text"`
		} `json:"languages"`
	} `json:"data"`
}

func GetUserStats() (*UserStats, error) {
	/// No. You dont wanna se my api key shared at this repo.
	apiKey := ""
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://wakatime.com/api/v1/users/current/stats/last_7_days", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Basic "+apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("some thing does wrong with wakatime :(")
	}
	var stats UserStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
