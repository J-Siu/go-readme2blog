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
	"bufio"
	"os"
	"strings"

	"github.com/J-Siu/go-helper"
)

type FileCutter struct {
	Bottom      *[]string `json:"bottom"`      // out: content after split marker
	Filename    string    `json:"filename"`    // in: filename
	NoSkip      bool      `json:"noSkip"`      // in: ignore skip-marker
	Skipped     bool      `json:"skip"`        // out: true if skip-marker found
	SkipMarker  string    `json:"skipMarker"`  // in: skip marker
	Split       bool      `json:"split"`       // out: true if split-marker is found
	SplitMarker string    `json:"splitMarker"` // in: split marker
	SplitLine   int       `json:"splitLine"`   // out: Line number split-marker is found
	Top         *[]string `json:"top"`         // out: content before split marker
}

type FileCutterList []*FileCutter

// Use <bottom>.Bottom
func (self *FileCutter) Join(bottom *FileCutter) *FileCutter {
	self.Bottom = bottom.Bottom
	return self
}

// Read file and split content base on parameters
//   - If <filename> empty, self.Filename will be used
//   - If readTop, clear self.Top, save content before split-marker to self.Top
//   - If readBottom, clear self.Bottom, save content after split-marker to self.Bottom
//   - Set self.Split = true if split-marker found, set self.SplitLn
//   - Set self.Skip = true if skip-marker found, clear self.SplitLn, self.Found
func (self *FileCutter) ReadContent(filename string, readTop bool, readBottom bool) *FileCutter {
	var err error = nil
	var file *os.File
	var bottom, top []string

	// Reset split and splitLine
	self.Split = false
	self.SplitLine = 0

	// Content is reset first
	if readBottom {
		self.Bottom = nil
	}
	if readTop {
		self.Top = nil
	}

	// Change markers to lower case
	self.SkipMarker = strings.ToLower(self.SkipMarker)
	self.SplitMarker = strings.ToLower(self.SplitMarker)

	if filename == "" {
		filename = self.Filename
	} else {
		self.Filename = filename
	}

	file, err = os.Open(filename)
	defer file.Close()

	if err == nil && !helper.IsRegularFile(filename) {
		err = helper.Err(filename + " is not a regular file.")
	}

	if err == nil {

		// Read file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			if !self.NoSkip && strings.ToLower(scanner.Text()) == self.SkipMarker {
				self.Reset()
				self.Skipped = true
				break
			}

			if !self.Split && strings.ToLower(scanner.Text()) == self.SplitMarker {
				self.Split = true
				continue // not including splitter in Bottom
			}

			if self.Split {
				if readBottom {
					bottom = append(bottom, scanner.Text())
				} else {
					break
				}
			} else {
				self.SplitLine++
				if readTop {
					top = append(top, scanner.Text())
				}
			}
		}
		// Finish reading file
		if self.Split {
			// Only accept content if split-marker found
			self.SplitLine += 1
			if readBottom {
				self.Bottom = &bottom
			}
			if readTop {
				self.Top = &top
			}
		} else {
			// no split-marker, reset line counter
			self.SplitLine = 0
		}
	}

	if err == nil && !self.Skipped && !self.Split {
		err = helper.Err(filename + ": No splitter found.")
	}

	if err != nil {
		helper.Errs.Add(err)
	}

	return self
}

// Read file and split content
//   - If <filename> empty, self.Filename will be used
//   - clear self.Top, save content before split-marker to self.Top
//   - clear self.Bottom, save content after split-marker to self.Bottom
//   - Set self.Split = true if split-marker found, set self.SplitLn
//   - Set self.Skip = true if skip-marker found, clear self.SplitLn, self.Found
func (self *FileCutter) Read(filename string) *FileCutter {
	return self.ReadContent(filename, true, true)
}

// Read file and split content
//   - If <filename> empty, self.Filename will be used
//   - clear self.Bottom, save content after split-marker to self.Bottom
//   - Set self.Split = true if split-marker found, set self.SplitLn
//   - Set self.Skip = true if skip-marker found, clear self.SplitLn, self.Found
func (self *FileCutter) ReadBottom(filename string) *FileCutter {
	return self.ReadContent(filename, false, true)
}

// Read file and split content
//   - If <filename> empty, self.Filename will be used
//   - clear self.Top, save content before split-marker to self.Top
//   - Set self.Split = true if split-marker found, set self.SplitLn
//   - Set self.Skip = true if skip-marker found, clear self.SplitLn, self.Found
func (self *FileCutter) ReadTop(filename string) *FileCutter {
	return self.ReadContent(filename, true, false)
}

// Reset all values except markers
func (self *FileCutter) Reset() *FileCutter {
	self.Bottom = nil
	self.Top = nil
	self.Skipped = false
	self.Split = false
	self.SplitLine = 0
	return self
}

// Save split content to file
//   - If <filename> empty, self.Filename will be used
//   - Save self.Top, split-maker, self.Bottom
//   - Not save if self.Top or self.Bottom is nil
func (self *FileCutter) Save(filename string) *FileCutter {
	var err error
	var file *os.File

	if filename == "" {
		filename = self.Filename
	} else {
		self.Filename = filename
	}

	if self.Bottom == nil || self.Top == nil {
		err = helper.Err(filename + " not all parts ready to be saved.")
	}

	// Open file
	if err == nil {
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		defer file.Close()
	}

	// Write file
	if err == nil {
		dataWriter := bufio.NewWriter(file)
		defer dataWriter.Flush()

		// Write top
		for _, data := range *self.Top {
			_, err = dataWriter.WriteString(data + "\n")
			if err != nil {
				break
			}
		}

		// Write splitter
		_, err = dataWriter.WriteString(self.SplitMarker + "\n")

		// Write bottom
		if err == nil {
			for _, data := range *self.Bottom {
				_, err = dataWriter.WriteString(data + "\n")
				if err != nil {
					break
				}
			}
		}
	}

	if err != nil {
		helper.Errs.Add(err)
	}

	return self
}

func (self *FileCutterList) GetNames() *[]string {
	var names []string
	for _, i := range *self {
		names = append(names, i.Filename)
	}
	return &names
}
