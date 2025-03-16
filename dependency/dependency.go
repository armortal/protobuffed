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

// package dependency provides an interface for installing a dependency.
package dependency

import (
	"context"

	"github.com/armortal/protobuffed/cache"
)

// Dependency represents a dependency that can be installed. It is up to the
// implementer to define how the dependency is installed.
type Dependency interface {

	// Install installs the dependency with the given version at the cache directory. All implementations
	// should assume the dir provided is empty.
	Install(ctx context.Context, dir *cache.Directory, version string) error
}
