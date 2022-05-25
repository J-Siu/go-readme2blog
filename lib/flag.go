/*
Copyright © 2022 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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

// Holding all flags from command line
type TypeFlag struct {
	Args          []string // Args from command line
	Debug         bool     // Enable debug output
	DefaultMdExt  string   // Default md extension
	DefaultReadme string   // Default readme filename
	DirBlog       string   // Hugo blog content dir
	DirOut        string   // Output directory
	DirSrc        string   // Parent directory of repositories(not dir repo itself)
	FileBlog      string   // Single blog file to be process
	FileOut       string   // Output file
	FileReadme    string   // Single readme to be process
	Forced        bool     // Allow overwriting original file
	MarkerSkip    string   // skip marker
	MarkerSplit   string   // split marker
	NoError       bool     // Do not print error
	NoParallel    bool     // Do not process in parallel(go routine)
	NoSkip        bool     // Flag for ignoring skip marker
	ShowFileList  bool     // Show file list in directory mode
}
