package main

import (
	"github.com/bansatya/v4/app"
)

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
