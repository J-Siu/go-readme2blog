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
	"path"
	"strconv"

	"github.com/J-Siu/go-helper"
)

var (
	Conf TypeConf
	Flag TypeFlag
)

const Version = "v1.0.2"

const (
	DEFAULT_MD_EXT       = ".md"
	DEFAULT_README       = "README.md"
	DEFAULT_MARKER_SKIP  = "<!--skip-sync-->"
	DEFAULT_MARKER_SPLIT = "<!--more-->"
	TXT_IN_SAME_DIR      = " in same directory."
	TXT_SAME_DIR         = " are same directories."
	TXT_NOT_DIR          = " is not directory."
	TXT_NOT_FILE         = " is not a regular file."
)

func init() {
	Flag.DefaultMdExt = DEFAULT_MD_EXT
	Flag.DefaultReadme = DEFAULT_README
	Flag.MarkerSkip = DEFAULT_MARKER_SKIP
	Flag.MarkerSplit = DEFAULT_MARKER_SPLIT
}

// Helper functions that are Flag dependent

func CheckMarker(listSkip, listSplit *FileCutterList, filename string) {
	fileCutter := FileCutterNew(filename).ReadTop("")
	fileCutter.Top = nil
	if fileCutter.Skipped {
		*listSkip = append(*listSkip, fileCutter)
	} else if fileCutter.Split {
		*listSplit = append(*listSplit, fileCutter)
	}
}

// Create map[blog file full path](repository readme full path)
//   - map is created if simplified blog name == simplified repository name
func DirFileMapInit() {
	Conf.Blog = make(FileMap)
	Conf.Blog.MapFile(Flag.DirBlog)
	Conf.Readme = make(FileMap)
	Conf.Readme.MapDirFile(Flag.DirSrc, Flag.DefaultReadme)
	Conf.ReadmeBlog = *Conf.Readme.Join(&Conf.Blog)
	if Flag.Debug {
		ReadmeBlogMapPrint()
	}
}

// Print readme - blog mapping
func ReadmeBlogMapPrint() {
	if Flag.Debug || Flag.ShowFileList {
		helper.Report(&Conf.Blog, "Blog List", true, false)
		helper.Report(&Conf.Readme, "Readme Folder List", true, false)
	}
	helper.Report(&Conf.ReadmeBlog, "Readme --> Blog ("+strconv.Itoa(len(Conf.ReadmeBlog))+")", false, false)
}

// Create *FileCutter
//   - Set markers using Flag
func FileCutterNew(filename string) *FileCutter {
	var self FileCutter
	self.Filename = filename
	self.NoSkip = Flag.NoSkip
	self.SkipMarker = Flag.MarkerSkip
	self.SplitMarker = Flag.MarkerSplit
	return &self
}

// map in format map[<readme>]=<blog>
func MappedFileSync() {
	var fileCutter *FileCutter = FileCutterNew("")
	for readme, blog := range Conf.ReadmeBlog {
		fileCutter.ReadTop(blog).ReadBottom(readme)
		if fileCutter.Bottom != nil || fileCutter.Top != nil {
			fileCutter.Save(path.Join(Flag.DirOut, path.Base(blog)))
			helper.Report(readme+":"+blog, "Processed", true, true)
		} else {
			helper.Report(readme+":"+blog, "-Skipped-", true, true)
		}
		fileCutter.Reset()
	}
}
