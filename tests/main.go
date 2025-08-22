package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/goodylabs/easycli"
	"github.com/goodylabs/easycli/providers/github"
	"github.com/joho/godotenv"
)

func main() {
	devPath := filepath.Join(".development")

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	githubUser := os.Getenv("GITHUB_USER")
	githubRepo := os.Getenv("GITHUB_REPO")

	app := easycli.ConfigureGithubApp(&github.GithubOpts{
		User: githubUser,
		Repo: githubRepo,
	})
	if err := app.Run(devPath); err != nil {
		log.Fatal(err)
	}
}
