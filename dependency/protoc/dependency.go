package protoc

import (
	"context"
	"fmt"
	"runtime"

	"github.com/armortal/protobuffed"
	"github.com/armortal/protobuffed/cache"
	"github.com/armortal/protobuffed/util"
)

type Dependency struct{}

func (d *Dependency) Install(ctx context.Context, dir *cache.Directory, version string) error {
	// Get the platform.
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "win64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch_64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "darwin":
		platform = "osx-universal_binary"
	default:
		return protobuffed.ErrRuntimeNotSupported
	}

	release := fmt.Sprintf("protoc-%s-%s.zip", version, platform)

	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/%s/%s", version, release)
	if err := util.Download(url, dir.Path()); err != nil {
		return err
	}

	if err := util.ExtractZip(release, dir.Path()); err != nil {
		return err
	}

	return nil
}

// func protobufVersionPath(cache string, version string) string {
// 	return filepath.Join(cache, "protobuf", version)
// }

// func protobufBinaryName() string {
// 	f := "protoc"
// 	if runtime.GOOS == "windows" {
// 		f += ".exe"
// 	}
// 	return f
// }
// func protobufBinaryPath(cache string, version string) string {
// 	return filepath.Join(protobufVersionPath(cache, version), "bin", protobufBinaryName())
// }
