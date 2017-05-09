package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"github.com/karantin2020/cli"
)

func TestArgs(t *testing.T) {
	flags := cli.New("", "")
	flags.Parse()
	flags.Usage()
}

func TestUsage(t *testing.T) {
	flags := cli.New("", "")
	base_app := filepath.Base(os.Args[0])
	flags.Parse()
	flags.Usage = func() {
		fmt.Printf(`Usage: %s -t tmpl -o outfile -pkg outpkg [--tag=value]...
Type '%s -h' | '%s --help' to see that message
Type '%s -t tmpl -tags' to see all tags in template
  -t: template to execute
  -o: what file to write
  -pkg: output package
`, base_app, base_app, base_app, base_app)
	}
	flags.Usage()
}