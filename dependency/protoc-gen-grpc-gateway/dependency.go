package protocgengrpcgateway

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/armortal/protobuffed"
	"github.com/armortal/protobuffed/cache"
	"github.com/armortal/protobuffed/util"
)

type Dependency struct{}

func (d *Dependency) Install(ctx context.Context, dir *cache.Directory, version string) error {
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "windows-x86_64.exe"
		case "arm64":
			platform = "windows-arm64.exe"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-arm64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "darwin-x86_64"
		case "arm64":
			platform = "darwin-arm64"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	default:
		return protobuffed.ErrRuntimeNotSupported
	}

	release := fmt.Sprintf("protoc-gen-grpc-gateway-%s-%s", version, platform)

	// /v2.19.1/protoc-gen-grpc-gateway-v2.19.1-darwin-arm64
	url := fmt.Sprintf("https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v%s/%s", version, release)
	fmt.Println(url)
	bin := filepath.Join(dir.Path(), "bin")
	// Create the bin folder
	if err := os.MkdirAll(bin, 0700); err != nil {
		return err
	}

	output := filepath.Join(bin, "protoc-gen-grpc-gateway")
	if runtime.GOOS == "windows" {
		output += ".exe"
	}

	if err := util.Download(url, output); err != nil {
		return err
	}

	return nil
}
