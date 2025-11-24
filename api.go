package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadRepos(language string, amount int) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	baseUrl := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&sort=stars&order=desc&per_page=%d", language, amount)

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("User-Agent", "GitHub-CLI")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	repos, err := parseJson(data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, repo := range repos {
		fmt.Printf("Repository:\n Name: %s\n Description: %s\n Stars: %d\n Language: %s\n URL: %s\n\n",
			repo.Name, repo.Description, repo.StarCount, repo.Language, repo.HTMLURL)
	}

}
