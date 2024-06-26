# Chalet 🏘️

Welcome to Chalet! Chalet provides a streamlined way to set up and manage development environments
using containers. With Chalet, you can run all your development tools, languages, and applications
in isolated containers, eliminating the need to install any languages or tools directly on your local machine.

* **Language Agnostic**: Supports a wide range of programming languages including Python, Node.js, Ruby, Java, and more.
* **Isolated Environments**: Each project runs in its own container, ensuring clean and conflict-free development setups.
* **Easy Setup** : Spin up development environments with a single command.
* **Consistent Environments**: Ensure all team members are using the same development setup, reducing "it works on my machine" issues.
* **Portable**: Easily share your development environment configurations with others.
  
## Usage
```
Provides a streamlined way to set up and manage development
environments using Docker containers.

With Chalet, you can run all your development tools, languages, and applications
in isolated containers, eliminating the need to install any languages or tools directly
on your local machine.

Usage:
  chalet [command]

Available Commands:
  exec        Execute your custom commands
  help        Help about any command
  init        Initialize a new chalet project
  run         Run the dev command for your project

Flags:
  -h, --help      help for chalet
  -v, --verbose   verbose logging

Use "chalet [command] --help" for more information about a command.
```

### Commands
- `chalet init` Initializes a new chalet project by creating a chalet.yml file,
which will contain the configuration for the project.

|   Flags                  |  Description            |
| ------------------------ | ----------------------- |
|  -h, --help              | help for init           |
|  -l, --language          | project language        |
|  -n, --name              | container project name  |
|  -p, --port              | server port             |
|  -r, --run               | run command             |
|      --version           | project version         |

- `chalet exec` The exec command allows you to run custom commands defined in your configuration file.
These commands are designed to simplify and automate various tasks within your project.

Examples:

1. Running a custom command defined in your chalet.yml configuration file:
  `chalet exec my_custom_command`

2. Executing an arbitrary shell command within the Chalet container:
  `chalet exec "echo Hello, World!`

The command first checks if the provided command exists in the custom commands defined in your
configuration file (chalet.yml). If it exists, it executes the corresponding command.
If not, it treats the input as a regular shell command and executes it within the container.

- `chalet run` Used to run the server locally on the chalet container, defined by the run command on the config file.

### Configuration
The configuration is set on the `chalet.yml` config file

| Configuration   | Type   | Description                                | Example              | required |
|-----------------|--------|--------------------------------------------|----------------------|----------|
| name            | string | project name                               | example              | yes      |
| language        | string | project language (docker image name)       | node                 | yes      |
| version         | string | language version                           | latest               | no       |
| server_port     | number | port where the server runs                 | 8080                 | yes      |
| exposed_port    | number | port exposed by chalet                     | 8081                 | no       |
| commands.run    | string | command used to run the server in dev mode | "npm run dev"        | yes      |
| custom_commands | object | map with custom and common commanda        | install: npm install | no       |


## Contributing
We welcome contributions! Please see our [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to get started.

## Issues and Feedback
If you encounter any issues or have feedback, please open an issue on GitHub. We appreciate your input and will do our best to address any problems.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Acknowledgements
We'd like to thank all the contributors and the open-source community for their invaluable support and contributions.

