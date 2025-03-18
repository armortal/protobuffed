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

package custom

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"strings"

	"github.com/armortal/protobuffed"
	"github.com/armortal/protobuffed/cache"
)

type Dependency struct{}

func (d *Dependency) Install(ctx context.Context, dir *cache.Directory, value string) error {
	u, err := url.Parse(value)
	if err != nil {
		return err
	}
	switch u.Scheme {
	case "http", "https":
		return d.installHttp(ctx, dir, u)
	case "git":
		return d.installGit(ctx, dir, u)
	default:
		return protobuffed.ErrDependencyNotSupported
	}
}

func (d *Dependency) installHttp(ctx context.Context, dir *cache.Directory, url *url.URL) error {
	r, err := http.Get(url.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := dir.Write(path.Base(url.Path), b, 0700); err != nil {
		return err
	}

	return nil
}

func (d *Dependency) installGit(ctx context.Context, dir *cache.Directory, url *url.URL) error {
	u := strings.Replace(url.String(), "git://", "https://", 1)
	cmd := exec.Command("sh", "-c", fmt.Sprintf("git clone %s %s", u, dir.Path()))
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
