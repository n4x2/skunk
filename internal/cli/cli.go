// Package cli provides a simple structure command-line application.
package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/n4x2/zoo/command"
)

const (
	ExitSuccess = iota
	ExitFailure
	ExitUsage
)

// App represents a command-line application.
type App struct {
	Banner   string             // Banner for the application.
	Name     string             // Name of the application.
	Brief    string             // Brief description of the application.
	Usage    string             // Usage information for the application.
	Version  string             // Version information for the application.
	Commands []*command.Command // List of available commands.
}

// AddCommand adds a new command to the application. It takes the 'c' type of [Cmd].
// It returns an error if a command with the same name already exists.
func (app *App) AddCommand(c *command.Command) error {
	for _, cmd := range app.Commands {
		if cmd.N == c.N {
			return fmt.Errorf("Command %s already exist", c.N)
		}
	}

	app.Commands = append(app.Commands, c)

	return nil
}

// Exit exits the application with the given code.
func (app *App) Exit(code int) {
	os.Exit(code)
}

// Help prints the help information for the application.
func (app *App) Help() {
	fmt.Println(app.Banner)
	fmt.Println(app.Brief)
	fmt.Printf("\nUsage:\n  %s\n\n", app.Usage)
	if len(app.Commands) > 0 {
		fmt.Println("Commands:")
		for _, cmd := range app.Commands {
			fmt.Printf("  %s\t%s\n", cmd.N, cmd.B)
		}
		fmt.Print("\n")
	}
	fmt.Println("Flags:")
	fmt.Println("  -h, --help\tShow this text")
	fmt.Println("  -v, --version\tShow version information")
	fmt.Println("\nUse '[command] --help' to see more information about a command.")
}

// NewCommand creates a new command. It takes the name (n), brief description (b), usage
// information (u), and a function (h) that will be executed when the command is invoked.
// It returns a command.
func (app *App) NewCommand(n, b, u string, h command.HandlerFunc) *command.Command {
	return command.New(n, b, u, h)
}

// Run executes the application with the given arguments.
func (app *App) Run(args []string) (int, error) {
	flag.Usage = func() { app.Help() }

	if len(args) == 1 {
		flag.Usage()
		app.Exit(ExitUsage)
	}

	var (
		helpFlag, versionFlag bool
	)

	flag.BoolVar(&helpFlag, "help", false, "Show help")
	flag.BoolVar(&helpFlag, "h", false, "Show help")
	flag.BoolVar(&versionFlag, "version", false, "Show version information")
	flag.BoolVar(&versionFlag, "v", false, "Show version information")
	flag.Parse()

	if helpFlag {
		flag.Usage()
		app.Exit(ExitUsage)
	}

	if versionFlag {
		if app.Version != "" {
			fmt.Println(app.Version)
			app.Exit(ExitSuccess)
		}

		if buildinfo, ok := debug.ReadBuildInfo(); ok && buildinfo.Main.Version != "" {
			fmt.Println(buildinfo.Main.Version)
			app.Exit(ExitSuccess)
		}

		fmt.Println("Unknown")
		app.Exit(ExitSuccess)
	}

	var command = args[1]
	for _, cmd := range app.Commands {
		if cmd.N == command {
			if err := cmd.Fs.Parse(args[2:]); err != nil {
				return ExitFailure, fmt.Errorf("unable to parse: argument %s", args[2])
			}

			if err := cmd.Fn(cmd.Fs, args[2:]); err != nil {
				return ExitFailure, err
			}
		}
	}

	return ExitFailure, fmt.Errorf("Unknown command: %s", command)
}

// New creates a new App instance with the given name and brief description.
func New(n, b string) *App {
	return &App{
		Name:     n,
		Brief:    b,
		Commands: make([]*command.Command, 0),
	}
}
