# protobuffed [![test](https://github.com/armortal/protobuffed/actions/workflows/test.yml/badge.svg)](https://github.com/armortal/protobuffed/actions/workflows/test.yml)

Protocol buffers buffed up :muscle: A lightweight tool for managing your protobuf projects.

Protobuffed was originally developed to ease the workload on developers when working with projects that utilize protocol buffers.
The process involved in setting up protobuf and plugin binaries can be overwhelming and time consuming particularly when working in teams and ensuring each developer has the same protobuf setup. Protobuffed aims to solve this issue by using a single configuration file that sits in your project's repository and does the heavy lifting so that developers don't have to.

> :warning: This project is currently in active development and we may make changes (potentially breaking) as we gather feedback from early adopters until we get to the first major release.

## Contents

- [Installation](#installation)
- [Getting Started](#getting-started)
	- [Initializing a new project](#initializing-a-new-project)
	- [Adding proto files](#adding-proto-files)
	- [Plugins](#plugins)
- [Configuration](#configuration)
- [Plugins](#plugins)
- [Commands](#commands)
	- [init](#init)
	- [print](#print)
	- [install](#install)
	- [generate](#generate)
- [Cache](#cache)
- [Contributing](#contributing)

## Installation

Install with `go`:

`go install github.com/armortal/protobuffed@latest`

## Getting Started

### Initializing a new project

Protobuffed uses a [configuration](#configuration) file that describes the project's plugins and its associated configuration. You can initialize a new project by running `protobuffed init` (generally in the root folder of your project). You will see a newly created file named `protobuffed.json` (can be changed with the `-f` or `--file` flag).

```json
{
    "protobuf": {
        "version": "21.12"
    },
    "imports": [],
    "inputs": [],
    "plugins": []
}
```

### Adding proto files

Once you have defined your project's `.proto` files, they need to be added to the `inputs` array. Any imports that you use in your protobuf files also need to be added as `imports` so that the protobuf compiler knows where to look for these imports.

### Plugins




### Installing protobuf

With the standard configuration, running `protobuffed install` will only install protobuf. When you run this, a [cache](#cache) folder named `.protobuffed` will be created which all binaries will be installed at. You can change this default location by specifying the `-c` or `--cache` flag. After running `install`, your project should now look like:

```
|--- .protobuffed
|    |--- protobuf/
|    |    |--- 21.12/
|    |
|    protobuffed.json

```

You can add the `.protobuffed` folder to your `.gitignore` so that the binaries aren't committed to Git.

## Configuration

```json
{
	"protobuf" : {
		"version": "21.12"
	},
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
		},
		{
			"name" : "grpc-web",
			"version" : "1.4.2",
			"options" : "import_style=commonjs+dts,mode=grpcwebtext",
			"output" : "testing"
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
| `name` | **string** | The plugin name (supported plugins are `go`,`go-grpc`, `grpc-web`). |
| `version` | **string** | The plugin version. |
| `options` | **string** | A comma separated string of plugin options in the form of KEY=VALUE (e.g. `KEY1=VALUE1,KEY2=VALUE2`)
| `output` | **string** | The output path. |

## Plugins

Protobuffed supports the following official plugins:

| Name | Source |
| :--- | :--------- |
| **go** | [https://github.com/protocolbuffers/protobuf-go](https://github.com/protocolbuffers/protobuf-go) |
| **go-grpc** | [https://github.com/grpc/grpc-go](https://github.com/grpc/grpc-go) |
| **grpc-web** | [https://github.com/grpc/grpc-web](https://github.com/grpc/grpc-web) |
| **js** | [https://github.com/protocolbuffers/protobuf-javascript](https://github.com/protocolbuffers/protobuf-javascript) |


## Commands

### init

`init` initializes a new configuration file.

### print

### install

### generate

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



## Cache

By default, all binaries are stored in the `$HOME/.protobuffed` directory. You can override the default by specifying the `--cache` or `-c` flag when running commands.

The `protobuf` binaries are located at `protobuf/${VERSION}`.

All the plugin binaries are located at `plugins/${PLUGIN_NAME}/${VERSION}`.

## Contributing

See [Contributing](./CONTRIBUTING.md).
