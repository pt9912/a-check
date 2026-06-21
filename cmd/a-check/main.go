// Command a-check is the thin entrypoint (ARC-006); the CLI logic lives in
// internal/cli so it stays black-box testable.
package main

import (
	"os"

	"github.com/pt9912/a-check/internal/cli"
)

func main() { os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr)) }
