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
	"errors"
	"os"

	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/J-Siu/go-helper/v2/str"
)

type FileStitchProperty struct {
	FileBottom  *string
	FileOut     *string
	FileTop     *string
	MarkerSkip  *string
	MarkerSplit *string
}

type FileStitch struct {
	*basestruct.Base
	*FileStitchProperty
}

func (t *FileStitch) New(property *FileStitchProperty) *FileStitch {
	t.Base = new(basestruct.Base)
	t.MyType = "FileStitch"
	prefix := t.MyType + ".New"

	if property != nil {
		t.FileStitchProperty = property
		if t.FileOut == nil || *t.FileOut == "" {
			t.Err = errors.New("no file out")
		}
		t.Initialized = true
	} else {
		t.Err = errors.New("property is nil")
	}
	ezlog.Debug().N(prefix).Lm(t).Out()

	errs.Queue(prefix, t.Err)
	return t
}

func (t *FileStitch) Reset() *FileStitch {
	t.FileBottom = nil
	t.FileOut = nil
	t.FileTop = nil
	t.MarkerSkip = nil
	t.MarkerSplit = nil

	t.Base = new(basestruct.Base)
	t.MyType = "FileStitch"
	t.Initialized = false
	return t
}

func (t *FileStitch) Run() *FileStitch {
	prefix := t.MyType + ".stitch"
	if t.CheckErrInit(prefix) {
		var (
			contentBottom *[]string
			contentOut    []string
			contentTop    *[]string
		)
		// --- Read files
		if t.Err == nil {
			contentBottom, t.Err = file.ReadStrArray(*t.FileBottom)
		}
		if t.Err == nil {
			contentTop, t.Err = file.ReadStrArray(*t.FileTop)
		}
		// --- Stitch
		if t.Err == nil {
			// ezlog.Debug().Nn(prefix).Nn("contentBottom").M(contentBottom).Out()
			// ezlog.Debug().Nn(prefix).Nn("contentTop").M(contentTop).Out()
			// This is not efficient, but we are dealing with file size < M byte
			if str.ArrayContains(contentBottom, t.MarkerSkip, false) {
				ezlog.Debug().N(prefix).N("SKIP").N("Skip marker found").M(t.FileBottom).Out()
			} else if str.ArrayContains(contentTop, t.MarkerSkip, false) {
				ezlog.Debug().N(prefix).N("SKIP").N("Skip marker found").M(t.FileTop).Out()
			} else if !str.ArrayContains(contentBottom, t.MarkerSplit, false) {
				ezlog.Debug().N(prefix).N("SKIP").N("Split marker not found").M(t.FileBottom).Out()
			} else if !str.ArrayContains(contentTop, t.MarkerSplit, false) {
				ezlog.Debug().N(prefix).N("SKIP").N("Split marker not found").M(t.FileTop).Out()
			} else {
				// stich - top from top file
				for _, item := range *contentTop {
					if item != *t.MarkerSplit {
						contentOut = append(contentOut, item)
					} else {
						break
					}
				}
				// stich - bottom from bottom file
				var markerFound bool
				for _, item := range *contentBottom {
					if !markerFound && item == *t.MarkerSplit {
						markerFound = true
					}
					if markerFound {
						contentOut = append(contentOut, item)
					}
				}
				if len(contentOut) > 0 {
					ezlog.Debug().N(prefix).N("FileOutContent").M(contentOut).Out()
					t.Err = file.WriteStrArray(*t.FileOut, &contentOut, os.FileMode(0600))
				}
			}
		}
		errs.Queue(prefix, t.Err)
	}
	return t
}
