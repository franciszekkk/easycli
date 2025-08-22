package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goodylabs/easycli/ports"
	"github.com/goodylabs/easycli/providers/github"
	"github.com/goodylabs/easycli/release"
	"github.com/goodylabs/easycli/utils"
	"github.com/joho/godotenv"
)

type EasyCliInstance struct {
	release  *release.ReleaseCfg
	provider ports.Provider
}

func (e *EasyCliInstance) Run(appPath string) error {

	configPath := filepath.Join(appPath, "config.json")

	if !e.release.CheckNeedsCheck(configPath) {
		return nil
	}

	fmt.Println("Checking for updates...")

	newestRelease, err := e.provider.GetNewestReleaseName()
	if err != nil {
		return err
	}

	if err := e.provider.PerformUpdate(""); err != nil {
		return err
	}

	e.release.ReleaseName = newestRelease
	e.release.LastCheck = utils.GetCurrentDate()
	return e.release.WriteReleaseCfg(configPath, e.release)
}

func ConfigureGithubApp(opts *github.GithubOpts) *EasyCliInstance {
	return &EasyCliInstance{
		release:  release.NewReleaseCfg(),
		provider: github.NewGithubApp(opts),
	}
}

func main() {
	devPath := filepath.Join(".development")

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	githubUser := os.Getenv("GITHUB_USER")
	githubRepo := os.Getenv("GITHUB_REPO")

	app := ConfigureGithubApp(&github.GithubOpts{
		User: githubUser,
		Repo: githubRepo,
	})
	if err := app.Run(devPath); err != nil {
		log.Fatal(err)
	}
}
