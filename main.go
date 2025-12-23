package main

import (
	"fmt"
	"os"

	"github.com/JulianH99/clone/cmd"
	"github.com/JulianH99/clone/internal"
	"github.com/JulianH99/clone/internal/ui"
)

func main() {
	if err := internal.CheckGit(); err != nil {
		fmt.Println(ui.InContainer("Git is not installed"))
		os.Exit(1)
	}
	cmd.Execute()
}
