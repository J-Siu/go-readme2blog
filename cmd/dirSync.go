/*
Copyright © 2023 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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
package cmd

import (
	"github.com/J-Siu/go-helper"
	"github.com/J-Siu/go-readme2blog/lib"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var dirSyncCmd = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s"},
	Short:   "Sync blog folder with readme in project folders",
	Long: `Sync blog folder with readme in project folders.

- Mapping is created between blog filenames and project folder names.

- Split marker(` + lib.Flag.MarkerSplit + `): Matching pair will have blog top part above split marker and ` + lib.Flag.DefaultReadme + ` part below splitter join together and put into output folder with blog filename. The pair is skipped if split marker is not found in one of the file. Split marker need to be in its own line.

- Skip marker(` + lib.Flag.MarkerSkip + `): No sync is performed if skip marker is found in one of the file. Skip marker should be placed above split marker and in its own line.`,
	Run: func(cmd *cobra.Command, args []string) {
		if helper.SameDir(lib.Flag.DirOut, lib.Flag.DirBlog) && !lib.Flag.Forced {
			helper.Errs.Add(helper.Err(lib.Flag.DirOut + ", " + lib.Flag.DirBlog + lib.TXT_SAME_DIR))
		}
		if helper.SameDir(lib.Flag.DirOut, lib.Flag.DirSrc) && !lib.Flag.Forced {
			helper.Errs.Add(helper.Err(lib.Flag.DirOut + ", " + lib.Flag.DirSrc + lib.TXT_SAME_DIR))
		}
		if !helper.IsDir(lib.Flag.DirOut) {
			helper.Errs.Add(helper.Err(lib.Flag.DirOut + lib.TXT_NOT_DIR))
		}
		if helper.Errs.NotEmpty() {
			return
		}

		lib.DirFileMapInit()
		lib.MappedFileSync()
	},
}

func init() {
	dirCmd.AddCommand(dirSyncCmd)
	dirSyncCmd.Flags().StringVarP(&lib.Flag.DirBlog, "dir-blog", "b", "", "Markdown directory")
	dirSyncCmd.Flags().StringVarP(&lib.Flag.DirOut, "dir-out", "o", "", "Output directory")
	dirSyncCmd.Flags().StringVarP(&lib.Flag.DirSrc, "dir-src", "s", "", "Source directory")
	dirSyncCmd.MarkFlagRequired("dir-blog")
	dirSyncCmd.MarkFlagRequired("dir-out")
	dirSyncCmd.MarkFlagRequired("dir-src")
}
