# protobuffed [![test](https://github.com/armortal/protobuffed/actions/workflows/test.yml/badge.svg)](https://github.com/armortal/protobuffed/actions/workflows/test.yml)

Protocol buffers buffed up :muscle: Making it easier to work with protobuf files and binaries.

This project was originally developed to ease the burden when setting up projects that utilize protocol buffers. It can be time consuming and frustrating for developers to set up projects with the correct binaries and plugins, and keep up to date with new versions without having to constantly download and reinstall them.

`protobuffed` uses a single configuration file that resides each project repository. The only requirement is for all developers to have this binary installed and everything else is handled seamlessly.

## Contents

- [Getting Started](#getting-started)
	- [Installing](#installing)
	- [Commands](#commands)
- [Configuration](#configuration)
- [Cache](#cache)
- [Contributing](#contributing)

## Getting Started

### Installing

Install with `go`:

`go install github.com/armortal/protobuffed@latest`

### Commands

Once installed, you can view all available commands and flags with `protobuffed --help`.

```
Protocol buffers buffed up. Making it easier to work with protobuf files and binaries

Usage:
  protobuffed [command]

Available Commands:
  command     Print the executable command.
  completion  Generate the autocompletion script for the specified shell
  generate    Generate source code.
  help        Help about any command
  install     Install binaries

Flags:
  -c, --cache string   The directory where binaries will be installed and executed from. (default "$HOME/.protobuffed")
  -f, --file string    The configuration file (default "protobuffed.json")
  -h, --help           help for protobuffed

Use "protobuffed [command] --help" for more information about a command.
```

To install binaries and generate source code, run `protobuffed generate -f protobuffed.json`.

To print the command which can be executed manually, run `protobuffed command -f protobuffed.json`.

If you would like to only install the binaries (no source code generated), run `protobuffed install -f protobuffed.json`.

## Configuration

A single configuration file (default is `protobuffed.json`) should reside in the repository where you will be generating your code and should be committed to Git. This configuration file is used as input to `protobuffed` to determine the versions of the binaries to install and execute. The following is an example of what a configuration file looks like:

```json
{
	"version" : "21.12",
	"imports" : [
		"test"
	],
	"inputs" : [
		"test.proto"
	],
	"plugins" : [
		{
			"name" : "go",
			"version" : "1.28.1",
			"options" : "paths=source_relative",
			"output" : "test"
		},
		{
			"name" : "go-grpc",
			"version" : "1.52.3",
			"options" : "paths=source_relative",
			"output" : "test"
		}
	]
}
```

| Name | Type | Description |
| :--- | :--- | :---------- |
| `version` | **string** | The version of protobuf to use. |
| `imports` | **[]string** | Imports to include. |
| `inputs` | **[]string** | Proto files to generate source for. |
| `plugins` | **[][Plugin](#plugin)** | Plugins to include. |

### Plugin

| Name | Type | Description |
| :--- | :--- | :---------- |
| `name` | **string** | The plugin name (supported plugins are `go`,`go-grpc`). |
| `version` | **string** | The plugin version. |
| `options` | **string** | A comma separated string of plugin options in the form of KEY=VALUE (e.g. `KEY1=VALUE1,KEY2=VALUE2`)
| `output` | **string** | The output path. |

## Cache

By default, all binaries are stored in the `$HOME/.protobuffed` directory. You can override the default by specifying the `--cache` or `-c` flag when running commands.

The `protobuf` binaries are located at `protobuf/${VERSION}`.

All the plugin binaries are located at `plugins/${PLUGIN_NAME}/${VERSION}`.

## Contributing

See [Contributing](./CONTRIBUTING.md).
