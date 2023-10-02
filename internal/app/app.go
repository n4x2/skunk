package app

import (
	"github.com/n4x2/skunk/internal/cli"
)

const banner = `
        __                 __
  _____|  | ____ __  ____ |  | __
 /  ___/  |/ /  |  \/    \|  |/ /
 \___ \|    <|  |  /   |  \    <
/____  >__|_ \____/|___|  /__|_ \
     \/     \/          \/     \/
`

const (
	name  = "skunk"
	brief = "A boring password manager"
	usage = name + " [COMMAND|FLAG] [ARGS]"
)

func NewApp() *cli.App {
	app := cli.New(name, brief)
	app.Banner = banner
	app.Usage = usage

	var (
		cmdName, cmdBrief, cmdUsage string
	)

	// Command generate.
	cmdName = "generate"
	cmdBrief = "Generate random non-consecutive string as password"
	cmdUsage = cmdName + " [FLAG] [ARGS]"
	generate := app.NewCommand(cmdName, cmdBrief, cmdUsage, generateHandler)

	// Command generate flags.
	generate.Fs.Bool("copy", false, "Copy generated password into clipboard")
	generate.Fs.Bool("numbers", false, `Include random numbers "0-9"`)
	generate.Fs.Bool("symbols", false, `Include random symbols "!@#$%^&*"`)
	generate.Fs.Int("len", 20, "The length of password")

	// Command generate examples.
	generate.E = map[string]string{
		"Copy into clipboard": "skunk generate --copy -numbers -symbols -len=18",
	}

	// Add commands into app.
	_ = app.AddCommand(generate)

	return app
}
