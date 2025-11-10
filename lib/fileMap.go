/*
MIT License

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package lib

import (
	"os"
	"path"

	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
)

type FileMap map[string]string

// Create map[simplified file name](file full path)
//   - create mapping between simplified base file names and full path of the file
func (t *FileMap) MapFile(filepath string) {
	(*t)[file.SimplifyName(filepath)] = filepath
}

// Create map[simplified file name](file full path) for all files in `dir`
//   - create mapping between simplified base file names and full path of the file
func (t *FileMap) MapFileDir(dir string) {
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		errs.Queue("", err)
		return
	}
	for _, f := range dirEntry {
		if f.Type().IsRegular() {
			(*t)[file.SimplifyName(f.Name())] = *file.FullPathStr(path.Join(dir, f.Name()))
		}
	}
}

// Create map[simplified dir name](file full path)
//   - create mapping between simplified dir names and full path of specific file it contains
func (t *FileMap) MapDirFile(dir, filename string) {
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		errs.Queue("", err)
		return
	}
	var realName string
	for _, d := range dirEntry {
		if d.Type().IsDir() {
			// Get real filename inside dir
			realName = file.InDir(path.Join(dir, d.Name()), filename)
			if realName != "" {
				(*t)[file.SimplifyName(d.Name())] = *file.FullPathStr(path.Join(dir, d.Name(), realName))
			}
		}
	}
}

// If self[<name>] and map2[<name>] exist, map3[self[<name>]] = map2[<name>].
//   - If an index exist in both maps, create a mapping with their values
//   - Return pointer of new map3
func (t *FileMap) Join(map2 *FileMap) *FileMap {
	var map3 = make(FileMap)

	for name1, value1 := range *t {
		if value2, exist := (*map2)[name1]; exist {
			map3[value1] = value2
		}
	}

	return &map3
}

func (t *FileMap) Dump() *FileMap {
	for k, v := range *t {
		ezlog.Log().N(k).M(v).Out()
	}
	return t
}
