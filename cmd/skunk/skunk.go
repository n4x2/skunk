package main

import (
	"fmt"
	"os"

	"github.com/n4x2/skunk/internal/app"
)

func main() {
	app := app.NewApp()
	if code, err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		app.Exit(code)
	}
}
