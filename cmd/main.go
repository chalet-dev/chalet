package main

import (
	"flag"
	"fmt"
	initproject "github.com/dfd/cli/init_project"
)

func main() {
    commandPtr := flag.String("command", "", "Command to execute.")
    flag.Parse()

    switch *commandPtr {
        case "init":
            initproject.InitProject()
        case "run":
            fmt.Println("Executing run command...")
        case "install":
            fmt.Println("Executing install command...")
        case "":
            fmt.Println("Command is required")
        default:
            fmt.Println("Unknown command:", *commandPtr)
    }
}