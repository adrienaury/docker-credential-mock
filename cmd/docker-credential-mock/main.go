package main

import (
	"fmt"
	"os"

	"github.com/adrienaury/docker-credential-mock/internal"
	"github.com/docker/docker-credential-helpers/credentials"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version   string
	commit    string
	buildDate string
	builtBy   string
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("%v (commit=%v date=%v by=%v)\n", version, commit, buildDate, builtBy)
		os.Exit(0)
	}
	credentials.Serve(internal.YAMLStorage{})
}
