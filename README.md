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
	- [Adding plugins](#adding-plugins)
	- [Generating code](#generating-code)
- [Configuration](#configuration)
- [Commands](#commands)
	- [generate](#generate)
	- [init](#init)
	- [install](#install)
	- [print](#print)
- [Cache](#cache)
- [Contributing](#contributing)

## Installation

Install with `go`:

`go install github.com/armortal/protobuffed@latest`

## Getting Started

This getting started guide is based off our example which you can find in the [examples directory](./examples/).

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

Each one of your projects will have at least one `.proto` file which will have service and message definitions. Let's create service and message definitions for an *auth* service in a file named `example.proto`.

```proto
syntax = "proto3";

option go_package = "github.com/armortal/protobuffed/examples/go;example";

package armortal.protobuffed.example;

service Auth {
    rpc SignIn(SignInRequest) returns (SignInResponse);

    rpc SignUp(SignUpRequest) returns (SignUpResponse);
}

message SignInRequest {
    string email = 1;
    string password = 2;
}

message SignInResponse {
    string token = 1;
}

message SignUpRequest {
    string email = 1;
    string password = 2;
}

message SignUpResponse {
    string token = 1;
}
```

Once you have defined your project's `.proto` files, they need to be added to the `inputs` array in the configuration file.

```json
{
    "protobuf": {
        "version": "21.12"
    },
    "imports": [],
    "inputs": [
        "example.proto"
    ],
    "plugins": []
}
```

> :warning: Any imports that you use in your protobuf files also need to be added to `imports` so that the protobuf compiler knows where to look for these imports.

Your project layout should now look like:

```
|--- example.proto
|--- protobuffed.json
```

### Adding plugins

You can one or more plugins in a single configuration file for your project. Protobuffed supports the following plugins:

| Name | Source |
| :--- | :--------- |
| **go** | [https://github.com/protocolbuffers/protobuf-go](https://github.com/protocolbuffers/protobuf-go) |
| **go-grpc** | [https://github.com/grpc/grpc-go](https://github.com/grpc/grpc-go) |
| **grpc-web** | [https://github.com/grpc/grpc-web](https://github.com/grpc/grpc-web) |
| **js** | [https://github.com/protocolbuffers/protobuf-javascript](https://github.com/protocolbuffers/protobuf-javascript) |

Let's create some folders where our generated code will be for `go` and `web` projects. Your project should now look like

```
|--- go/
|--- web/
|--- example.proto
|--- protobuffed.json
```

We'll now add the plugins to the configuration file:

```json
{
    "protobuf": {
        "version": "21.12"
    },
    "imports": [],
    "inputs": [
        "example.proto"
    ],
    "plugins": [
        {
            "name": "go",
            "version": "1.28.1",
            "options": "paths=source_relative",
            "output": "go/"
        },
        {
            "name": "go-grpc",
            "version": "1.52.3",
            "options": "paths=source_relative",
            "output": "go/"
        },
        {
            "name": "grpc-web",
            "version": "1.4.2",
            "options": "import_style=commonjs+dts,mode=grpcwebtext",
            "output": "web/"
        },
        {
            "name": "js",
            "version": "3.21.2",
            "options": "import_style=commonjs,binary",
            "output": "web/"
        }
    ]
}
```

### Generating code

Now that our configuration file is defined, we can now generate source code with `protobuffed generate`. When executing this command, Protobuffed will first check to see if the binaries exist in the [cache](#cache) and if not they will be installed first before the source code is generated.

> :information_source: If you would like to just install the binaries and not generate the source code, run `protobuffed install`.

If the default cache location is used, you will see a newly created folder named `.protobuffed` which contains `protoc` and `protoc-gen` binaries. You should update your `.gitignore` to include `.protobuffed/` so that the binaries aren't committed to Git. With the binaries installed and the code generated, your project should now look like:

```
├── .protobuffed/
|   ├── plugins/
|   ├── protobuf/
├── go/
|   ├── example.pb.go
|   ├── example_grpc.pb.go
├── web/
|   ├── example_grpc_web_pb.d.ts
|   ├── example_grpc_web_pb.js
|   ├── example_pb.d.ts
|   ├── example_pb.js
├── .gitignore
├── example.proto
├── protobuffed.json
```

## Configuration

A configuration file represents your project's protobuf and plugin configuration.

```json
{
    "protobuf": {
        "version": "21.12"
    },
    "imports": [],
    "inputs": [
        "example.proto"
    ],
    "plugins": [
        {
            "name": "go",
            "version": "1.28.1",
            "options": "paths=source_relative",
            "output": "go/"
        },
        {
            "name": "go-grpc",
            "version": "1.52.3",
            "options": "paths=source_relative",
            "output": "go/"
        },
        {
            "name": "grpc-web",
            "version": "1.4.2",
            "options": "import_style=commonjs+dts,mode=grpcwebtext",
            "output": "web/"
        },
        {
            "name": "js",
            "version": "3.21.2",
            "options": "import_style=commonjs,binary",
            "output": "web/"
        }
    ]
}
```

| Name | Type | Description |
| :--- | :--- | :---------- |
| `protobuf` | [Protobuf](#protobuf) | Protobuf configuration. |
| `imports` | **[]string** | Imports to include. |
| `inputs` | **[]string** | Proto files to generate source for. |
| `plugins` | **[][Plugin](#plugin)** | Plugins to include. |

### Protobuf

| Name | Type | Description |
| :--- | :--- | :---------- | 
| `version` | **string** | The version of protobuf to use. |

### Plugin

| Name | Type | Description |
| :--- | :--- | :---------- |
| `name` | **string** | The plugin name (supported plugins are `go`,`go-grpc`, `grpc-web`, and `js`). |
| `version` | **string** | The plugin version. |
| `options` | **string** | A comma separated string of plugin options in the form of KEY=VALUE (e.g. `KEY1=VALUE1,KEY2=VALUE2`)
| `output` | **string** | The output path. |

## Commands

All commands support the following flags:

- `-c` or `--cache` - path of the cache where binaries will be installed and executed from (default is `.protobuffed`)
- `-f` or `--file` - path of the configuration file (default is `protobuffed.json`)

### generate

`protobuffed generated` will generate source code for a given configuration. If the binaries are not installed, this command will install them first.

### init

`protobuffed init` initializes a new configuration file.

### install

`protobuffed install` will install all binaries for a given configuration.

### print

`protobuffed print` will print out the `protoc` executable command for a given configuration which you can use to run manually.

## Cache

By default, all binaries are stored in a `.protobuffed` directory in the folder where Protobuffed is executed. You can override the default by specifying the `-c` or `--cache` flag when running commands.

The `protobuf` binaries are located at `.protobuffed/protobuf/${VERSION}`.

The plugin binaries are located at `.protobuffed/plugins/${NAME}/${VERSION}`.

## Contributing

See [Contributing](./CONTRIBUTING.md).
