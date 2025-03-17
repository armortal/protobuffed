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

package util

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Extract extracts the archive at src to the directory dst. If the archive
// format is not supported or there was an error extracting, an error is returned.
func Extract(src string, dst string) error {
	switch filepath.Ext(src) {
	case ".tar":
		return ExtractTar(src, dst)
	case ".zip":
		return ExtractZip(src, dst)
	default:
		return fmt.Errorf("unsupported archive format")
	}
}

func ExtractTar(src string, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()
	return extractTarFromReader(f, dst)
}

func extractTarFromReader(r io.Reader, dst string) error {
	t := tar.NewReader(r)

	for {
		header, err := t.Next()

		// EOF means we are done.
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(dst, header.Name), 0700); err != nil {
				return err
			}
		case tar.TypeReg:
			out, err := os.Create(filepath.Join(dst, header.Name))
			if err := out.Chmod(0700); err != nil {
				return err
			}
			if err != nil {
				return err
			}
			if _, err := io.Copy(out, t); err != nil {
				return err
			}
			out.Close()

		default:
			return fmt.Errorf("unknown type %s", string(header.Typeflag))
		}
	}
}

func ExtractZip(src string, dst string) error {
	// Create the reader
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, f := range reader.File {
		if err := extractZip(f, dst); err != nil {
			return err
		}
	}
	return nil
}

func ExtractZipFromBytes(b []byte, dst string) error {
	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return err
	}

	for _, f := range r.File {
		if err := extractZip(f, dst); err != nil {
			return err
		}
	}
	return nil
}

func extractZip(f *zip.File, dst string) error {
	fp := filepath.Join(dst, f.Name)
	if !strings.HasPrefix(fp, filepath.Clean(dst)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", fp)
	}

	// If this is a directory, we need to create it.
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(fp, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(fp), os.ModePerm); err != nil {
		return err
	}

	out, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(out, zippedFile); err != nil {
		return err
	}
	return nil
}
