// The config package provides a common global store for configuration options.
package config

import (
	"fmt"
	"github.com/rkcpi/vell/repos"
	"github.com/rkcpi/vell/rpm"
	"os"
)

var (
	HttpPort    = os.Getenv("VELL_HTTP_PORT")
	HttpAddress = os.Getenv("VELL_HTTP_ADDRESS")
	ReposPath   = os.Getenv("VELL_REPOS_PATH")

	RepoStore     repos.RepositoryStore
	ListenAddress string
)

func init() {
	if HttpPort == "" {
		HttpPort = "8080"
	}

	if ReposPath == "" {
		ReposPath = "/var/lib/vell/repositories"
	}
	RepoStore = rpm.NewRepositoryStore(ReposPath)

	ListenAddress = fmt.Sprintf("%s:%s", HttpAddress, HttpPort)
}
