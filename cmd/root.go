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
	"os"

	"github.com/J-Siu/go-helper"
	"github.com/J-Siu/go-readme2blog/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "go-sync-readme-blog",
	Short:   "Sync Blog with README.md",
	Version: lib.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		helper.Debug = lib.Flag.Debug
		helper.ReportDebug(&lib.Flag, "Flag", false, false)
		lib.Conf.Init()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if !lib.Flag.NoError {
			helper.Report(helper.Warns, "Warning", true, false)
			helper.Report(helper.Errs, "Error", true, false)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&lib.Flag.Debug, "debug", "d", false, "Enable debug output")
	rootCmd.PersistentFlags().BoolVarP(&lib.Flag.Forced, "force", "F", false, "Enable overwriting original file")
	rootCmd.PersistentFlags().BoolVarP(&lib.Flag.NoError, "no-error", "", false, "Do not print error")
	rootCmd.PersistentFlags().BoolVarP(&lib.Flag.NoParallel, "no-parallel", "n", false, "Do not process in parallel")
	rootCmd.PersistentFlags().BoolVarP(&lib.Flag.NoSkip, "no-skip", "", false, "Ignore skip marker")
	rootCmd.PersistentFlags().BoolVarP(&lib.Flag.ShowFileList, "show-files", "l", false, "Show file lists")
	rootCmd.PersistentFlags().StringVarP(&lib.Conf.FileConf, "config", "", lib.DefaultConfFile, "Config file")
	rootCmd.PersistentFlags().StringVarP(&lib.Flag.DefaultMdExt, "md-ext", "", lib.DEFAULT_MD_EXT, "Markdown extension")
	rootCmd.PersistentFlags().StringVarP(&lib.Flag.DefaultReadme, "readme", "", lib.DEFAULT_README, "Readme filename")
	rootCmd.PersistentFlags().StringVarP(&lib.Flag.MarkerSkip, "skip-marker", "", lib.DEFAULT_MARKER_SKIP, "")
	rootCmd.PersistentFlags().StringVarP(&lib.Flag.MarkerSplit, "split-marker", "", lib.DEFAULT_MARKER_SPLIT, "")
}

func initConfig() {
	viper.SetConfigType("json")
	if lib.Conf.FileConf == "" {
		lib.Conf.FileConf = lib.DefaultConfFile
	}
	viper.SetConfigFile(helper.TildeEnvExpand(lib.Conf.FileConf))
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err != nil {
		helper.Report(err.Error(), "", true, true)
		os.Exit(1)
	}
}
