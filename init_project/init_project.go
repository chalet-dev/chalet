package initproject

import (
    "fmt"
    "os"
    "io"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Lang    string `yaml:"lang"`
    Version string `yaml:"version"`
}

func InitProject() {
    // Get the current working directory
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current directory:", err)
        return
    }

    // Open the YAML file
    file, err := os.Open(cwd + "/chalet.yaml")
    if err != nil {
        fmt.Println("Error opening YAML file:", err)
        return
    }
    defer file.Close()

    // Read the YAML file
    yamlFile, err := io.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading YAML file:", err)
        return
    }

    // Parse the YAML file into our configuration struct
    var config Config
    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        fmt.Println("Error parsing YAML:", err)
        return
    }

    // Print the values
    fmt.Println("Language:", config.Lang)
    fmt.Println("Version:", config.Version)
}
