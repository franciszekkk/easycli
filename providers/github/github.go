package github

import (
	"fmt"

	"github.com/goodylabs/easycli/adapters/httpconnector"
	"github.com/goodylabs/easycli/adapters/oshelper"
	"github.com/goodylabs/easycli/ports"
)

type GithubOpts struct {
	User string
	Repo string
}

type githubApp struct {
	opts           GithubOpts
	newReleaseName string
	newReleaseUrl  string
	httpconnector  *httpconnector.HttpClient
	oshelper       *oshelper.OsHelper
}

func NewGithubApp(opts *GithubOpts) ports.Provider {
	return &githubApp{
		opts:          *opts,
		httpconnector: httpconnector.NewHttpClient(),
	}
}

func (g *githubApp) GetNewestReleaseName() (string, error) {
	lastestReleaseUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", g.opts.User, g.opts.Repo)
	if err := g.httpconnector.DoGet(lastestReleaseUrl, &releaseRes); err != nil {
		return "", err
	}

	fmt.Println("Latest Release URL:", lastestReleaseUrl)

	g.newReleaseName = releaseRes.TagName

	osType := g.oshelper.GetOSType()
	osArch, err := g.oshelper.GetArch()
	if err != nil {
		return "", err
	}

	for _, asset := range releaseRes.Assets {
		if asset.Name == fmt.Sprintf("%s-%s", osType, osArch) {
			g.newReleaseUrl = asset.BrowserDownloadURL
			break
		}
	}

	return g.newReleaseName, nil
}

func (g *githubApp) PerformUpdate(binaryPath string) error {
	osType := g.oshelper.GetOSType()
	osArch, err := g.oshelper.GetArch()
	if err != nil {
		return err
	}

	if g.newReleaseUrl == "" {
		return fmt.Errorf("no compatible binary found for %s-%s", osType, osArch)
	}

	fmt.Println("Downloading binary from:", g.newReleaseUrl)

	if err := g.oshelper.DownloadBinary(g.newReleaseUrl, binaryPath); err != nil {
		return err
	}

	fmt.Println("Updated binary at:", binaryPath)

	return nil
}
