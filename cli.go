package cli

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/spf13/pflag"
)

// App contains application info needed to parse flags
type App struct {
	*pflag.FlagSet
	// App description
	Description string
	// App version
	Version string

	help bool
	ver bool
}

// Create new App instance
func New(desc, ver string) *App {
	app := &App{
		FlagSet: pflag.NewFlagSet(os.Args[0], pflag.ExitOnError),
		Description: desc,
		Version: ver,
	}
	app.Usage = func() {
		if app.Description != "" {
			fmt.Println(app.Description + "\n")
		}
		fmt.Fprintf(os.Stderr, "Usage: %s [<flags>] [<args>]\n", filepath.Base(os.Args[0]))
		app.PrintDefaults()
		fmt.Println("")
	}
	app.SortFlags = false
	app.BoolVarP(&app.help, "help", "h", false, "this help message")
	app.BoolVarP(&app.ver, "version", "v", false, "print this app version")
	return app
}

// Helper function to print version App info
func (this *App) PrintVersion() {
	fmt.Println(this.Version)
}

// Parse parses flags
func (this *App) Parse() {
	this.FlagSet.Parse(os.Args[1:])
	if this.help {
		this.Usage()
		os.Exit(1)
	}
	if this.ver {
		this.PrintVersion()
		os.Exit(1)
	}
}