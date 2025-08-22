package github

import (
	"fmt"

	"github.com/goodylabs/easycli/ports"
)

type githubApp struct {
	repoUrl string
}

func NewGithubApp(repoUrl string) ports.Provider {
	return &githubApp{
		repoUrl: repoUrl,
	}
}

func (g *githubApp) CheckForUpdates() (bool, error) {
	return false, fmt.Errorf("func CheckForUpdates - not implemented")
}

func (g *githubApp) PerformUpdate(path string) error {
	return fmt.Errorf("func PerformUpdate - not implemented")
}
