/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// trendingCmd represents the trending command
var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "Retrieve trending GitHub repositories",
	Run: func(cmd *cobra.Command, args []string) {
		language, _ := cmd.Flags().GetString("language")
		limit, _ := cmd.Flags().GetInt("limit")
		loadRepos(language, limit)

	},
}

func init() {
	trendingCmd.Flags().StringP("language", "l", "go", "Programming language to filter repositories")
	trendingCmd.Flags().IntP("limit", "n", 10, "Number of repositories to retrieve")
	rootCmd.AddCommand(trendingCmd)
}
