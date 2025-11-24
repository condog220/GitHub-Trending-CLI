package main

import (
	"fmt"
)

func main() {
	fmt.Println("GitHub Trending CLI")
	fmt.Print("Enter programming language: ")
	var language string
	fmt.Scanln(&language)
	fmt.Print("Enter number of repositories to fetch: ")
	var amount int
	fmt.Scanln(&amount)
	loadRepos(language, amount)
}
