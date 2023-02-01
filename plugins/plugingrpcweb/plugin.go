package plugingrpcweb

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/armortal/protobuffed/util"
)

type Plugin struct{}

func New() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Install(version string, dst string) error {
	release, err := release(version)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://github.com/grpc/grpc-web/releases/download/%s/%s", version, release)
	fmt.Println(url)
	bin := filepath.Join(dst, "bin")
	// Create the bin folder
	if err := os.MkdirAll(bin, 0700); err != nil {
		return err
	}

	output := filepath.Join(bin, "protoc-gen-grpc-web")
	if runtime.GOOS == "windows" {
		output += ".exe"
	}

	if err := util.Download(url, output); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) Name() string {
	return "grpc-web"
}

func release(version string) (string, error) {
	var platform string
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			platform = "windows-x86_64.exe"
		case "arm64":
			platform = "windows-aarch64.exe"
		default:
			return "", errRuntimeNotSupported(version)
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x86_64"
		case "arm64":
			platform = "linux-aarch64"
		default:
			return "", errRuntimeNotSupported(version)
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			platform = "darwin-x86_64"
		case "arm64":
			platform = "darwin-aarch64"
		default:
			return "", errRuntimeNotSupported(version)
		}
	default:
		return "", errRuntimeNotSupported(version)
	}

	return fmt.Sprintf("protoc-gen-grpc-web-%s-%s", version, platform), nil
}

func errRuntimeNotSupported(version string) error {
	return fmt.Errorf("runtime %s:%s not supported for plugin %s@%s", runtime.GOOS, runtime.GOARCH, "go", version)
}
