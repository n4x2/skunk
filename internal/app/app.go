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

	return app
}
