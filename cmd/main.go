package main

import (
    "flag"
    "fmt"
)

func main() {
    commandPtr := flag.String("command", "", "Command to execute.")
    flag.Parse()

    switch *commandPtr {
        case "init":
            fmt.Println("Executing init command...")
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