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
package cmd

import (
	"path"

	"github.com/J-Siu/go-helper"
	"github.com/J-Siu/go-readme2blog/lib"
	"github.com/spf13/cobra"
)

// fileSyncCmd represents the file command
var fileSyncCmd = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s"},
	Short:   "Sync blog file with readme file",
	Long: `Sync blog file with readme file.

- If output is directory, blog filename will be used.

- Split marker(` + lib.Flag.MarkerSplit + `): Blog file top part above split marker and readme file below splitter join together and put into output file. The pair is skipped if split marker is not found in one of the file. Split marker need to be in its own line.

- Skip marker(` + lib.Flag.MarkerSkip + `): No sync is performed if skip marker is found in one of the file. Skip marker need to be placed above split marker and in its own line.`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileOut string
		if helper.IsDir(lib.Flag.FileOut) {
			fileOut = path.Join(lib.Flag.FileOut, path.Base(lib.Flag.FileBlog))
		} else {
			fileOut = lib.Flag.FileOut
		}
		helper.Report(fileOut, "fileOut", false, true)
		if helper.SameDir(fileOut, lib.Flag.FileBlog) && !lib.Flag.Forced {
			helper.Errs.Add(helper.Err(lib.Flag.FileOut + ", " + lib.Flag.FileBlog + lib.TXT_IN_SAME_DIR))
		}
		if helper.SameDir(fileOut, lib.Flag.FileReadme) && !lib.Flag.Forced {
			helper.Errs.Add(helper.Err(lib.Flag.FileOut + ", " + lib.Flag.FileReadme + lib.TXT_IN_SAME_DIR))
		}
		if helper.Errs.NotEmpty() {
			return
		}

		var newBlog *lib.FileCutter = lib.FileCutterNew("").ReadTop(lib.Flag.FileBlog).ReadBottom(lib.Flag.FileReadme)
		if helper.Errs.Empty() && !newBlog.Skipped {
			newBlog.Save(fileOut)
		}
	},
}

func init() {
	fileSyncCmd.Flags().StringVarP(&lib.Flag.FileOut, "out", "o", "", "Output file or directory")
	fileSyncCmd.Flags().StringVarP(&lib.Flag.FileBlog, "blog", "b", "", "Blog file")
	fileSyncCmd.Flags().StringVarP(&lib.Flag.FileReadme, "readme", "r", "", "Readme file")
	fileSyncCmd.MarkFlagRequired("blog")
	fileSyncCmd.MarkFlagRequired("out")
	fileSyncCmd.MarkFlagRequired("readme")
	fileCmd.AddCommand(fileSyncCmd)
}
