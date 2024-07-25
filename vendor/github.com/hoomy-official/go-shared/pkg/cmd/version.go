package cmd

import (
	"fmt"

	"github.com/hoomy-official/go-shared/pkg/buildinfo"
)

// Version holds detailed build information for an application, including the
// name, the specific version string, commit hash, the source of the build, and the build date.
type Version struct {
	buildinfo.BuildInfo
}

// NewVersion creates a new Version instance with the given version details.
func NewVersion(name, version, commit, buildSource, date string) Version {
	return Version{buildinfo.NewBuildInfo(name, version, commit, buildSource, date)}
}

// Run outputs the version information to the console. It implements the interface
// required to integrate with the kong CLI parsing library.
func (u Version) Run() error {
	//nolint:forbidigo // using a custom writer is not necessary here
	fmt.Print(u.BuildInfo.String())
	return nil
}
