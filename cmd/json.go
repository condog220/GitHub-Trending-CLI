package cmd

import "encoding/json"

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StarCount   int    `json:"stargazers_count"`
	Language    string `json:"language"`
	HTMLURL     string `json:"html_url"`
}

type Response struct {
	Items []Repo `json:"items"`
}

func parseJson(data []byte) ([]Repo, error) {
	var response Response
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.Items, nil

}
