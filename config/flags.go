package config

import (
	"fmt"
	"os"
)

type options struct {
	Version func() `long:"version" description:"Show version"`
}

// Options app flags
var Options options

var (
	// BuildTime build time variables
	BuildTime string

	// CommitHash build time variables
	CommitHash string
)

func init() {
	Options.Version = func() {
		fmt.Printf("Build Time: %s\n", CommitHash)
		fmt.Printf("Version: %s\n", BuildTime)
		os.Exit(0)
	}
}
