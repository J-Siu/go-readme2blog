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

// fileChkCmd represents the fileChk command
var fileChkCmd = &cobra.Command{
	Use:     "check <file ...>",
	Aliases: []string{"c"},
	Short:   "Check for skip and split marker",
	Run: func(cmd *cobra.Command, args []string) {
		var listSkip, listSplit lib.FileCutterList
		for _, f := range args {
			lib.CheckMarker(&listSkip, &listSplit, f)
		}
		helper.Report(listSkip.GetNames(), "Have skip marker", true, false)
		helper.Report(listSplit.GetNames(), "Have split marker", true, false)
	},
}

func init() {
	fileCmd.AddCommand(fileChkCmd)
}
