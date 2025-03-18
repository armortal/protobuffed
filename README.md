# protobuffed [![test](https://github.com/armortal/protobuffed/actions/workflows/test.yml/badge.svg)](https://github.com/armortal/protobuffed/actions/workflows/test.yml)

Protocol buffers buffed up :muscle: A lightweight tool for managing your protobuf projects.

Protobuffed was originally developed to ease the workload on developers when working with projects that utilize protocol buffers. The process involved in setting up protobuf and plugin binaries can be overwhelming and time consuming particularly when working in teams and ensuring each developer have the same binary versions. Protobuffed aims to solve this issue by using a single configuration file that sits in your project's repository and does the heavy lifting so that developers don't have to.

> :warning: This project is currently in active development and we may make changes (potentially breaking) as we gather feedback from early adopters until we get to the first major release.

## Contents

- [Installation](#installation)
- [Getting Started](#getting-started)
	- [Initializing a new project](#initializing-a-new-project)
    - [Creating your proto](#creating-your-proto)
	- [Adding your configuration](#adding-your-configuration)
	- [Generating code](#generating-code)
- [Dependencies](#dependencies)
    - [Registered](#registered)
    - [Custom](#custom)
- [Configuration](#configuration)
- [Commands](#commands)
	- [init](#init)
	- [install](#install)
  - [run](#run)
  - [generate](#generate)
- [Cache](#cache)
- [Contributing](#contributing)

## Installation

Install with `go`:

`go install github.com/armortal/protobuffed`

## Getting Started

This getting started guide is based off our example which you can find in the [examples directory](./examples).

### Initializing a new project

Protobuffed uses a [configuration](#configuration) file that describes the project's dependencies, plugins and its associated configuration. You can initialize a new project by running `protobuffed init` (generally in the root folder of your project). You will see a newly created file named `protobuffed.json` (can be changed with the `-f` or `--file` flag).

The below example was created with `protobuffed init --name example`

```json
{
  "name": "example",
  "dependencies": {
    "protoc": "v30.1"
  },
  "imports": [],
  "inputs": [],
  "plugins": [],
  "scripts" : {}
}
```

### Creating your proto

Each one of your projects will have at least one `.proto` file which will have service and message definitions. Let's create service and message definitions for an *Auth* service in a file named `example.proto`. In our example, we are using the third
party [Google APIs](https://github.com/googleapis/googleapis) options for generating our gRPC Gateway stub.

```proto
syntax = "proto3";

option go_package = "github.com/armortal/protobuffed/examples";

import "google/api/annotations.proto";

package armortal.protobuffed.example;

service Auth {
    rpc SignIn(SignInRequest) returns (SignInResponse) {
        option (google.api.http) = {
            post: "/signin"
            body: "*"
        };
    };

    rpc SignUp(SignUpRequest) returns (SignUpResponse) {
        option (google.api.http) = {
            post: "/signup"
            body: "*"
        };
    };
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

We now need to add the following to our configuration:
- The name of the proto file in the **inputs** array.
- The plugins and their associated dependencies.
- The Google APIs **dependency** so we can import the third party proto definitions.
- The imports of both the Google APIs dependency and the location of our protos (we need to add this if specifying an import).

```json
{
  "name": "example",
  "dependencies": {
    "protoc": "v30.1",
    "protoc-gen-go": "v1.36.5",
    "protoc-gen-go-grpc": "v1.71.0",
    "protoc-gen-grpc-gateway": "v2.26.1",
    "protoc-gen-grpc-web": "v1.5.0",
    "protoc-gen-js": "v3.21.4",
    "googleapis": "git://github.com/googleapis/googleapis"
  },
  "imports": [
    ".",
    ".protobuffed/googleapis"
  ],
  "inputs": [
    "example.proto"
  ],
  "plugins": [
    {
      "name": "go",
      "options": "paths=source_relative",
      "output": "./"
    },
    {
      "name": "go-grpc",
      "options": "paths=source_relative",
      "output": "./"
    },
    {
      "name": "grpc-gateway",
      "options": "paths=source_relative",
      "output": "./"
    },
    {
      "name": "grpc-web",
      "options": "import_style=commonjs+dts,mode=grpcwebtext",
      "output": "./"
    },
    {
      "name": "js",
      "options": "import_style=commonjs,binary",
      "output": "./"
    }
  ]
}
```

### Install dependencies

In order to generate our code, we need to first run `protobuffed install` to download and install our dependencies. If your
configuration file location is different to the default (`protobuffed.json`), you can specify this with the `--file` of `-f` option.

After running the command, you will see a newly created folder named `.protobuffed` which contains all your dependencies. 
You should update your `.gitignore` to include `.protobuffed/` so that the dependencies aren't committed to Git. With the binaries installed and the code generated, your project should now look like:

```
├── .protobuffed/
|   ├── protoc/
|   ├── protoc-gen-go/
|   ├── protoc-gen-go-grpc/
|   ├── protoc-gen-grpc-gateway/
|   ├── protoc-gen-grpc-web/
|   ├── protoc-gen-js/
|   ├── googleapis
├── .gitignore
├── example.proto
├── protobuffed.json
```

### Generating code

Now that our configuration file is defined and our dependencies installed, we can now generate source code with `protobuffed generate`. Below is an example of your directory after executing the command.

> :information_source: By default, each dependencies' **bin** folder is added to the path before executing **protoc**.

```
├── .protobuffed/
|   ├── protoc/
|   ├── protoc-gen-go/
|   ├── protoc/
|   ├── protoc-gen-go/
|   ├── protoc-gen-go-grpc/
|   ├── protoc-gen-grpc-gateway/
|   ├── protoc-gen-grpc-web/
|   ├── protoc-gen-js/
|   ├── googleapis
├── example.pb.go
├── example_grpc.pb.go
├── example_grpc_web_pb.d.ts
├── example_grpc_web_pb.js
├── example_pb.d.ts
├── example_pb.js
├── .gitignore
├── example.proto
├── protobuffed.json
```

## Dependencies

All projects will have at least one dependency. There are both [registered](#registered) and [custom](#custom) dependencies which can be included.

### Registered

The following dependencies are registered and implemented in this project. They can be used directly with semantic versions (e.g. v0.1.0)

| Name | Source |
| :--- | :--------- |
| **protoc** | [https://github.com/protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf) |
| **protoc-gen-go** | [https://github.com/protocolbuffers/protobuf-go](https://github.com/protocolbuffers/protobuf-go) |
| **protoc-gen-go-grpc** | [https://github.com/grpc/grpc-go](https://github.com/grpc/grpc-go) |
| **protoc-gen-grpc-gateway** | [https://github.com/grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) |
| **protoc-gen-grpc-web** | [https://github.com/grpc/grpc-web](https://github.com/grpc/grpc-web) |
| **protoc-gen-js** | [https://github.com/protocolbuffers/protobuf-javascript](https://github.com/protocolbuffers/protobuf-javascript) |

### Custom

If you require a dependency that is not registered, you can include them using a URL with the following schemes:

| Scheme | Description |
| :----- | :---------- |
| `git` | Clone a git repository using `git clone`. You must have `git` installed. |
| `http` | Download via http. |
| `https` | Download via https. |

If you need to add some further logic than just simply downloading a custom dependency (e.g. build a binary, unzip a file), you can add
it to the `scripts` section of your configuration file and execute it with the `run` command prior to generating your files.

For example, to continue from our example above, if we need to download this repository's contents and unzip it we can add an inline script or call one like the following:

```json
{
  "name": "example",
  "dependencies": {
    "protoc": "v30.1",
    "protoc-gen-go": "v1.36.5",
    "protoc-gen-go-grpc": "v1.71.0",
    "protoc-gen-grpc-gateway": "v2.26.1",
    "protoc-gen-grpc-web": "v1.5.0",
    "protoc-gen-js": "v3.21.4",
    "googleapis": "git://github.com/googleapis/googleapis",
    "armortal": "https://github.com/armortal/protobuffed/archive/refs/heads/main.zip"
  },
  "imports": [
    ".",
    ".protobuffed/googleapis"
  ],
  "inputs": [
    "example.proto"
  ],
  "plugins": [
    {
      "name": "go",
      "options": "paths=source_relative",
      "output": "./"
    },
    {
      "name": "go-grpc",
      "options": "paths=source_relative",
      "output": "./"
    },
    {
      "name": "grpc-gateway",
      "options": "paths=source_relative",
      "output": "./"
    },
    {
      "name": "grpc-web",
      "options": "import_style=commonjs+dts,mode=grpcwebtext",
      "output": "./"
    },
    {
      "name": "js",
      "options": "import_style=commonjs,binary",
      "output": "./"
    }
  ],
  "scripts": {
    "unzipInline": "unzip -o .protobuffed/armortal/main.zip -d .protobuffed/armortal && mv .protobuffed/armortal/protobuffed-main/* .protobuffed/armortal && rm -rf .protobuffed/armortal/protobuffed-main .protobuffed/armortal/main.zip",
    "unzip": "scripts/unzip.sh"
  }
}
```

## Configuration

A configuration file represents your project's configuration.

| Name | Type | Description |
| :--- | :--- | :---------- |
| `dependencies` | **map[string]string** | The dependency configuration. If dependency value is a semantic version (e.g. v1.0.0), this must be a registered dependency defined in this project. If not, you must use a `http(s)` or `git` URL. |
| `imports` | **[]string** | Imports to include. |
| `inputs` | **[]string** | Proto files to generate source for. |
| `plugins` | **[][Plugin](#plugin)** | Plugins to include. |
| `scripts` | **map[string]string** | A map of scripts to define that can be executed via the `run` command. |

### Plugin

| Name | Type | Description |
| :--- | :--- | :---------- |
| `name` | **string** | The plugin name. This must be the full binary name and must be found in a dependencies' **bin** folder or environment **PATH**. |
| `options` | **string** | A comma separated string of plugin options in the form of KEY=VALUE (e.g. `KEY1=VALUE1,KEY2=VALUE2`)
| `output` | **string** | The output path. |

## Commands

Execute commands with `protobuffed <COMMAND> <OPTIONS>`

| Name | Description |
| :--- | :---------- |
| [init](#init) | Initializes a new configuration file. |
| [install](#install) | Install all dependencies. |
| [run](#run) | Run a script in the configuration. |
| [generate](#generate) | Run the protoc compiler and generate source files. |

### init

Initializes a new configuration file.

| Options | Short | Description |
| :------ | :---- | :---------- |
| `file` | `f` | The path of the configuration file to write (default is `protobuffed.json`) |
| `name` | `n` | The name of the project. |

### install

Install all dependencies.

| Options | Short | Description |
| :------ | :---- | :---------- |
| `file` | `f` | The path of the configuration file to read (default is `protobuffed.json`) |


### run

Run a script defined in the configuration file. The command is `protobuffed run <SCRIPT_NAME>` where **SCRIPT_NAME** is
one that is defined in the `scripts` section of the configuration file.

| Options | Short | Description |
| :------ | :---- | :---------- |
| `file` | `f` | The path of the configuration file to read (default is `protobuffed.json`) |

### generate

Run the protoc compiler and generate source files.

| Options | Short | Description |
| :------ | :---- | :---------- |
| `file` | `f` | The path of the configuration file to read (default is `protobuffed.json`) |

## Cache

All dependencies are stored in a `.protobuffed` directory in the folder where Protobuffed is executed.

Each subfolder named after the key name in your dependencies' configuration.

For custom dependencies, if you want binaries to be included in the execution path, they must be located in
a `bin` folder within the dependency folder (e.g. `.protobuffed/<YOUR_DEPENDENCY>/bin`)

## Contributing

See [Contributing](./CONTRIBUTING.md).
