package protocgengo

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

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
			platform = "windows.amd64.zip"
		case "arm64":
			platform = "windows.arm64.zip"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux.amd64.tar.gz"
		case "arm64":
			platform = "linux.arm64.tar.gz"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "darwin.amd64.tar.gz"
		case "arm64":
			platform = "darwin.arm64.tar.gz"
		default:
			return protobuffed.ErrRuntimeNotSupported
		}
	default:
		return protobuffed.ErrRuntimeNotSupported
	}

	// Get the release and url.
	release := fmt.Sprintf("protoc-gen-go.%s.%s", version, platform)
	url := fmt.Sprintf("https://github.com/protocolbuffers/protobuf-go/releases/download/%s/%s", version, release)

	// Download and extract.
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := dir.Write(release, b, 0700); err != nil {
		return err
	}

	// Create the bin folder
	bin := filepath.Join(dir.Path(), "bin")
	if err := os.Mkdir(bin, 0700); err != nil {
		return err
	}

	if strings.HasSuffix(release, ".zip") {
		if err := util.ExtractZip(release, bin); err != nil {
			return err
		}
	} else {
		if err := util.ExtractTarGz(release, bin); err != nil {
			return err
		}
	}
	return nil
}
