package version

import "runtime"

type Version struct {
	GoVersion string
	GoOs      string
	GoArch    string
	GitRepo   string
	GitCommit string
	BuildDate string
	BuildUser string
	BuildTag  string
}

func Set(gitRepo, gitCommit, buildDate, buildUser, buildTag string) *Version {
	return &Version{
		GoVersion: runtime.Version(),
		GoOs:      runtime.GOOS,
		GoArch:    runtime.GOARCH,
		GitRepo:   gitRepo,
		GitCommit: gitCommit,
		BuildDate: buildDate,
		BuildUser: buildUser,
		BuildTag:  buildTag,
	}
}
