package github

import (
	"fmt"

	"github.com/goodylabs/easycli/adapters/httpconnector"
	"github.com/goodylabs/easycli/ports"
)

type GithubOpts struct {
	User string
	Repo string
}

type githubApp struct {
	opts                   GithubOpts
	newestAvailableRelease string
	httpconnector          *httpconnector.HttpClient
}

func NewGithubApp(opts *GithubOpts) ports.Provider {
	return &githubApp{
		opts:          *opts,
		httpconnector: httpconnector.NewHttpClient(),
	}
}

func (g *githubApp) GetNewestReleaseName() (string, error) {
	var releaseRes struct {
		TagName string `json:"tag_name"`
	}
	lastestReleaseUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", g.opts.User, g.opts.Repo)
	if err := g.httpconnector.DoGet(lastestReleaseUrl, &releaseRes); err != nil {
		return "", err
	}

	g.newestAvailableRelease = releaseRes.TagName

	return g.newestAvailableRelease, nil
}

func (g *githubApp) PerformUpdate(path string) error {
	return fmt.Errorf("func PerformUpdate - not implemented")
}

// os_type=$(uname -s | tr '[:upper:]' '[:lower:]')
// arch=$(uname -m)

// case "$arch" in
//     x86_64)
//         arch="amd64"
//         ;;
//     aarch64 | arm64)
//         arch="arm64"
//         ;;
//     *)
//         echo "Unsupported architecture: $arch"
//         exit 1
//         ;;
// esac

// artifact_url=$(curl -s "$releaseUrl" | jq -r ".assets[] | select(.name | test(\"tug-${os_type}-${arch}\")) | .browser_download_url")

// if [[ -z "$artifact_url" ]]; then
//     echo "No compatible binary found for ${os_type}-${arch}"
//     exit 1
// fi

// echo "Downloading Tug binary for ${os_type}-${arch}..."

// curl --fail --location --progress-bar --compressed --retry 3 --retry-delay 5 \
//   --max-time 10 -o "$TUG_BIN_PATH/tug" "$artifact_url"

// chmod +x "$TUG_BIN_PATH/tug"

// echo "Tug has been installed successfully!"
