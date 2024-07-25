package buildinfo

import "fmt"

// BuildInfo holds detailed build information for an application, including the
// name, the specific version string, commit hash, the source of the build, and the build date.
type BuildInfo struct {
	name        string
	version     string
	commit      string
	buildSource string
	date        string
}

// NewBuildInfo creates a new BuildInfo instance with the given version details.
func NewBuildInfo(name, version, commit, buildSource, date string) BuildInfo {
	return BuildInfo{name, version, commit, buildSource, date}
}

// Name returns the name of the application.
func (u BuildInfo) Name() string {
	return u.name
}

// Version returns the version string of the application.
func (u BuildInfo) Version() string {
	return u.version
}

// Commit returns the commit hash corresponding to this build of the application.
func (u BuildInfo) Commit() string {
	return u.commit
}

// BuildSource returns the source from which the application was built.
func (u BuildInfo) BuildSource() string {
	return u.buildSource
}

// Date returns the build date of the application.
func (u BuildInfo) Date() string {
	return u.date
}

func (u BuildInfo) String() string {
	return fmt.Sprintf(
		"name=%s, version=%s, commit=%s, buildDate=%s, buildSource=%s",
		u.Name(), u.Version(), u.Commit(), u.Date(), u.BuildSource(),
	)
}
