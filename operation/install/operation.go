// Copyright 2025 ARMORTAL TECHNOLOGIES PTY LTD

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 		http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package install

import (
	"context"

	"github.com/armortal/protobuffed"
	"github.com/armortal/protobuffed/cache"
	"github.com/armortal/protobuffed/config"
	"github.com/armortal/protobuffed/dependency"
	"github.com/armortal/protobuffed/dependency/custom"
	"github.com/armortal/protobuffed/dependency/protoc"
	protocgengo "github.com/armortal/protobuffed/dependency/protoc-gen-go"
	protocgengogrpc "github.com/armortal/protobuffed/dependency/protoc-gen-go-grpc"
	protocgengrpcgateway "github.com/armortal/protobuffed/dependency/protoc-gen-grpc-gateway"
	protocgengrpcweb "github.com/armortal/protobuffed/dependency/protoc-gen-grpc-web"
	protocgenjs "github.com/armortal/protobuffed/dependency/protoc-gen-js"
	"golang.org/x/mod/semver"
)

func Execute(ctx context.Context, cfg *config.Config, cache *cache.Cache) error {
	return new(Operation).Execute(ctx, cfg, cache)
}

type Operation struct{}

func (o *Operation) Execute(ctx context.Context, cfg *config.Config, cache *cache.Cache) error {
	if err := cache.RemoveAll(); err != nil {
		return err
	}

	// Install the dependencies
	for name, version := range cfg.Dependencies {
		var dep dependency.Dependency
		// If the version is not a semantic version, this is a custom dependency.
		if !semver.IsValid(version) {
			dep = new(custom.Dependency)

		} else {
			switch name {
			case "protoc":
				dep = new(protoc.Dependency)
			case "protoc-gen-go":
				dep = new(protocgengo.Dependency)
			case "protoc-gen-go-grpc":
				dep = new(protocgengogrpc.Dependency)
			case "protoc-gen-grpc-gateway":
				dep = new(protocgengrpcgateway.Dependency)
			case "protoc-gen-grpc-web":
				dep = new(protocgengrpcweb.Dependency)
			case "protoc-gen-js":
				dep = new(protocgenjs.Dependency)
			default:
				return protobuffed.ErrDependencyNotSupported
			}
		}

		// Install the dependency
		if err := dep.Install(ctx, cache.Directory(name), version); err != nil {
			return err
		}
	}

	return nil
}
