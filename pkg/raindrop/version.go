package raindrop

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/djherbis/times"
)

// Version parts (overridden at build time)
var (
	Major  = "0"
	Minor  = "5"
	Patch  = "7"
	Suffix = "" // empty by default, use -X to set dev during local builds
)

// Version returns a version string.
// prefix=true -> add "v" at start
// suffix=true -> add Suffix if non-empty
// versionFull=true -> add OS/ARCH and build timestamp
func Version(prefix, suffix, versionFull bool) string {
	version := fmt.Sprintf("%s.%s.%s", Major, Minor, Patch)

	if prefix {
		version = fmt.Sprintf("v%s", version)
	}

	if suffix && Suffix != "" {
		version = fmt.Sprintf("%s-%s", version, Suffix)
	}

	if versionFull {
		version = fmt.Sprintf("%s-%s-%s", version, runtime.GOOS, runtime.GOARCH)

		var creationTime time.Time
		path, err := exec.LookPath("raindrop")
		if err == nil && path != "" {
			t, err := times.Stat(path)
			if err == nil && t.HasBirthTime() {
				creationTime = t.BirthTime()
			}
		}

		if !creationTime.IsZero() {
			version = fmt.Sprintf("%s - built %v", version, creationTime.Format("2006-01-02 15:04:05"))
		}
	}

	return version
}
