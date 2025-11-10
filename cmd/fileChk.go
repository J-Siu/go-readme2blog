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
	"github.com/J-Siu/go-helper/v2/array"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-readme2blog/global"
	"github.com/spf13/cobra"
)

var fileChkCmd = &cobra.Command{
	Use:     "check <file ...>",
	Aliases: []string{"c"},
	Short:   "Check for skip and split marker",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			listNoMarker,
			listSkip,
			listSplit array.Array[string]
		)
		for _, f := range args {
			hasSkip, hasSplit, e := global.ChkMarker(f)
			if e == nil {
				if hasSkip {
					listSkip.Add(f)
				}
				if hasSplit {
					listSplit.Add(f)
				}
				if !hasSkip && !hasSplit {
					listNoMarker.Add(f)
				}
			}
		}
		ezlog.Log().Se().Lp(false).N("Skip").Lm(&listSkip).Out()
		ezlog.Log().Se().Lp(false).N("Split").Lm(&listSplit).Out()
		ezlog.Log().Se().Lp(false).N("No Marker").Lm(&listNoMarker).Out()
	},
}

func init() {
	fileCmd.AddCommand(fileChkCmd)
}
