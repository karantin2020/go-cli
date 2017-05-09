package main

import (
	// "fmt"
	// "os"
	// "path/filepath"
	"github.com/karantin2020/cli"
)

func main() {
	flags := cli.New("Test cli app", "0.0.1")
// 	base_app := filepath.Base(os.Args[0])
// 	flags.Usage = func() {
// 		fmt.Printf(`Usage: %s -t tmpl -o outfile -pkg outpkg [--tag=value]...
// Type '%s -h' | '%s --help' to see that message
// Type '%s -t tmpl -tags' to see all tags in template
//   -t: template to execute
//   -o: what file to write
//   -pkg: output package
// `, base_app, base_app, base_app, base_app)
// 	}
	flags.Parse()
	// flags.Usage()
	
	// flags.Usage()
}
