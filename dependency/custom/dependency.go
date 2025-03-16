package custom

import (
	"context"
	"net/url"

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
	return nil
}

func (d *Dependency) installGit(ctx context.Context, dir *cache.Directory, url *url.URL) error {
	return nil
}
