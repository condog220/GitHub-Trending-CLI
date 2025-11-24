# GitHub Trending CLI

A command-line interface (CLI) tool written in Go to fetch and display trending GitHub repositories based on stars.

## Features

-   **Fetch Trending Repos**: Get a list of popular repositories.
-   **Filter by Language**: Specify a programming language to filter the results (e.g., Go, Python, JavaScript).
-   **Custom Limit**: Control the number of repositories displayed.
-   **Output**: Displays repository name, description, star count, language, and URL.

## Prerequisites

-   [Go](https://go.dev/dl/) (version 1.25 or higher)
-   A GitHub Personal Access Token (for API authentication)

## Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/condog220/GitHub-Trending-CLI.git
    cd GitHub-Trending-CLI
    ```

## Configuration

This tool requires a GitHub API key to interact with the GitHub API.

1.  Create a `.env` file in the root directory of the project:

    ```bash
    touch .env
    ```

2.  Add your GitHub API key to the `.env` file:

    ```env
    API_KEY=your_github_personal_access_token
    ```

    > **Note:** You can generate a Personal Access Token in your GitHub [Developer Settings](https://github.com/settings/tokens).

## Usage

You can run the tool directly using `go run main.go` or build it into a binary.

### Basic Usage

Fetch the top 10 trending Go repositories (default):

```bash
go run main.go trending
```

### Filter by Language

Fetch trending repositories for a specific language (e.g., Python) using the `-l` or `--language` flag:

```bash
go run main.go trending -l python
```

### Set Limit

Specify the number of repositories to retrieve using the `-n` or `--limit` flag:

```bash
go run main.go trending -n 20
```

### Combine Flags

Fetch the top 5 JavaScript repositories:

```bash
go run main.go trending --language javascript --limit 5
```
