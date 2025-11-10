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
	"errors"

	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/J-Siu/go-readme2blog/global"
	"github.com/J-Siu/go-readme2blog/txt"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var dirSyncCmd = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s"},
	Short:   "Sync blog folder with readme in project folders",
	Long: `Sync blog folder with readme in project folders.

- Mapping is created between blog filenames and project folder names.

- Split marker(` + global.Conf.MarkerSplit + `): Matching pair will have blog top part above split marker and ` + global.DEFAULT_README + ` part below splitter join together and put into output folder with blog filename. The pair is skipped if split marker is not found in one of the file. Split marker need to be in its own line.

- Skip marker(` + global.Conf.MarkerSkip + `): No sync is performed if skip marker is found in one of the file. Skip marker should be placed above split marker and in its own line.`,
	Run: func(cmd *cobra.Command, args []string) {
		if file.SameDir(global.Flag.DirOut, global.Flag.DirBlog) && !global.Flag.Forced {
			errs.Queue("", errors.New(global.Flag.DirOut+", "+global.Flag.DirBlog+txt.SAME_DIR))
		}
		if file.SameDir(global.Flag.DirOut, global.Flag.DirSrc) && !global.Flag.Forced {
			errs.Queue("", errors.New(global.Flag.DirOut+", "+global.Flag.DirSrc+txt.SAME_DIR))
		}
		if !file.IsDir(global.Flag.DirOut) {
			errs.Queue("", errors.New(global.Flag.DirOut+txt.NOT_DIR))
		}
		if errs.NotEmpty() {
			return
		}

		mapReadmeBlog := global.MapReadmeBlog(global.Flag.DirBlog, global.Flag.DirSrc)
		global.MapReadmeBlogSync(mapReadmeBlog, global.Flag.DirOut)
	},
}

func init() {
	dirCmd.AddCommand(dirSyncCmd)
	dirSyncCmd.Flags().StringVarP(&global.Flag.DirBlog, "dir-blog", "b", "", "Markdown directory")
	dirSyncCmd.Flags().StringVarP(&global.Flag.DirOut, "dir-out", "o", "", "Output directory")
	dirSyncCmd.Flags().StringVarP(&global.Flag.DirSrc, "dir-src", "s", "", "Source directory")
	dirSyncCmd.MarkFlagRequired("dir-blog")
	dirSyncCmd.MarkFlagRequired("dir-out")
	dirSyncCmd.MarkFlagRequired("dir-src")
}
