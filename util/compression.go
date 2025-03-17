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
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Decompress(src string, dst string) error {
	switch filepath.Ext(src) {
	case ".gz":
		return DecompressGz(src, dst)
	default:
		return fmt.Errorf("unsupported compression format")
	}
}

func DecompressGz(src string, dst string) error {
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

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// Write the decompressed data to the destination file.
	if err := os.WriteFile(dst, b, 0700); err != nil {
		return err
	}

	return nil
}
