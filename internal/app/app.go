package app

import (
	"github.com/n4x2/skunk/internal/app/handler"
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

	// Command add.
	cmdName = "add"
	cmdBrief = "Add new password into vault"
	cmdUsage = "add [FLAG] [ARGS]"
	add := app.NewCommand(cmdName, cmdBrief, cmdUsage, handler.AddPassword)
	// Command add flags.
	add.Fs.String("name", "", "Password name")
	// Command add examples.
	add.E = map[string]string{
		"Password name": `skunk add --name="Streaming"`,
	}

	// Command edit.
	cmdName = "edit"
	cmdBrief = "Edit existing password in vault"
	cmdUsage = "edit [FLAG] [ARGS]"
	edit := app.NewCommand(cmdName, cmdBrief, cmdUsage, handler.EditPassword)
	// Command add flags.
	edit.Fs.String("name", "", "Password name")
	// Command add examples.
	edit.E = map[string]string{
		"Edit password": `skunk edit --name="Git"`,
	}

	// Command find.
	cmdName = "find"
	cmdBrief = "Find existing password by matching the name"
	cmdUsage = "find [FLAG] [ARGS]"
	find := app.NewCommand(cmdName, cmdBrief, cmdUsage, handler.FindPassword)
	//Commmand find flags.
	find.Fs.String("name", "", "Password name")
	// Command find examples.
	find.E = map[string]string{
		"Password name": `skunk find --name="Git"`,
	}

	// Command list.
	cmdName = "list"
	cmdBrief = "List available passwords"
	cmdUsage = ""
	list := app.NewCommand(cmdName, cmdBrief, cmdUsage, handler.ListPassword)
	list.Fs.String("pass", "", "Password for vault instead of using prompt")

	// Command generate.
	cmdName = "generate"
	cmdBrief = "Generate random non-consecutive string as password"
	cmdUsage = cmdName + " [FLAG] [ARGS]"
	generate := app.NewCommand(cmdName, cmdBrief, cmdUsage, handler.GeneratePassword)
	// Command generate flags.
	generate.Fs.Bool("copy", false, "Copy generated password into clipboard")
	generate.Fs.Bool("numbers", false, `Include random numbers "0-9"`)
	generate.Fs.Bool("symbols", false, `Include random symbols "!@#$%^&*"`)
	generate.Fs.Int("len", 20, "The length of password")
	// Command generate examples.
	generate.E = map[string]string{
		"Copy into clipboard": "skunk generate --copy -numbers -symbols -len=18",
	}

	// Command remove.
	cmdName = "rm"
	cmdBrief = "Remove password from vault"
	cmdUsage = "rm [FLAG] [ARGS]"
	remove := app.NewCommand(cmdName, cmdBrief, cmdUsage, handler.RemovePassword)
	//Commmand remove flags.
	remove.Fs.String("name", "", "Password name")
	// Command remove examples.
	remove.E = map[string]string{
		"Remove by name": `skunk rm --name="Git"`,
	}

	// Command show.
	cmdName = "show"
	cmdBrief = "Show existing password or copy into clipboard"
	cmdUsage = "show [FLAG] [ARGS]"
	show := app.NewCommand(cmdName, cmdBrief, cmdUsage, handler.ShowPassword)
	//Commmand show flags.
	show.Fs.String("name", "", "Password name")
	show.Fs.String("pass", "", "Password vault instead of using prompt")
	show.Fs.Bool("copy", false, "Copy password into clipboard")
	// Command show examples.
	show.E = map[string]string{
		"Copy into clipboard": `skunk show --name="Git" --copy`,
	}

	// Add commands into app.
	app.AddCommand(add, edit, find, list, generate, remove, show)

	return app
}
