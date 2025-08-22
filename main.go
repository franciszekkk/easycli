package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/goodylabs/easycli/ports"
	"github.com/goodylabs/easycli/providers/github"
	"github.com/goodylabs/easycli/release"
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

	// hasUpdate, err := e.app.CheckForUpdates()
	// if err != nil {
	// 	return err
	// } else if !hasUpdate {
	// 	return nil
	// }

	// return e.app.PerformUpdate("")

	// e.release.LastCheck = utils.GetCurrentDate()

	return e.release.WriteReleaseCfg(configPath, e.release)
}

func ConfigureGithubApp(githubUrl string) *EasyCliInstance {
	return &EasyCliInstance{
		release:  release.NewReleaseCfg(),
		provider: github.NewGithubApp(githubUrl),
	}
}

func main() {
	devPath := filepath.Join(".development")
	app := ConfigureGithubApp("")
	if err := app.Run(devPath); err != nil {
		log.Fatal(err)
	}
}
