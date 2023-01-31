# protobuffed [![test](https://github.com/armortal/protobuffed/actions/workflows/test.yml/badge.svg)](https://github.com/armortal/protobuffed/actions/workflows/test.yml)

Protocol buffers buffed up :muscle: Making it easier to work with protobuf files and binaries.

This project was originally developed to ease the burden when setting up projects with protobuf. Issues arise when working in teams and having each developer set up the same versions of the binaries. It can be time consuming and somewhat frustrating modifying paths, installing the right plugins, etc. and also difficult to keep up to date with the latest binaries.

**Protobuffed** uses a single configuration file that resides each project repository. The only requirement is for all developers to have this binary installed and everything else is handled seamlessly.

## Contents

- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Storage](#storage)
- [Contributing](#contributing)

## Getting Started

### Installing

`go install github.com/armortal/protobuffed@latest`

### Commands

To generate source, run `protobuffed generate -f protobuffed.json`. This will call **install** and then generate the source.

To install binaries only, run `protobuffed install -f protobuffed.json`.

To view the executable command, run `protobuffed print -f protobuffed.json`.

To view help, run `protobuffed --help`.

```
Protocol buffers buffed up. Making it easier to work with protobuf files and binaries

Usage:
  protobuffed [flags]
  protobuffed [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate source code
  help        Help about any command
  install     Install binaries
  print       Print the executable command.

Flags:
  -f, --file string   The configuration file (default "protobuffed.json")
  -h, --help          help for protobuffed

Use "protobuffed [command] --help" for more information about a command.
```

## Configuration

This project uses a configuration file (default is `protobuffed.json`). This configuration file should reside in the repository where you will be generating your code and should be committed to Git. 
The following is an example of what a configuration file looks like:

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
			"version" : "1.52.0",
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
| `name` | **string** | The plugin name. (`go`,`go-grpc`) |
| `version` | **string** | The plugin version. |
| `options` | **string** | A comma separated string of plugin options in the form of KEY=VALUE (e.g. `KEY1=VALUE1,KEY2=VALUE2`)
| `output` | **string** | The output path. |

## Storage

By default, all binaries are stored in the `$HOME/.protobuffed` directory. You can override the default by setting the environment variable `PROTOBUFFED_HOME`.

The `protobuf` binaries are located at `protobuf/VERSION`.

All the plugin binaries are located at `plugins/protoc-gen-${PLUGIN_NAME}/${VERSION}`.

## Contributing

See [Contributing](./CONTRIBUTING.md).
