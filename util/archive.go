// MIT License

// Copyright (c) 2023 Armortal Technologies Pty Ltd

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package util

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ExtractTarGz(src string, dst string) error {
	// Open the file to create a reader.
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	// Create a new gzip reader.
	r, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer r.Close()

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
