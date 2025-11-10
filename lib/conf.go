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
	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/spf13/viper"
)

var Default = TypeConf{
	FileConf: "~/.go-readme2blog.json",
}

type TypeConf struct {
	*basestruct.Base
	FileConf    string `json:"FileConf"`
	MarkerSkip  string `json:"MarkerSkip"`  // skip marker
	MarkerSplit string `json:"MarkerSplit"` // split marker
}

// Fill in conf struct from viper
func (t *TypeConf) New() {
	t.Base = new(basestruct.Base)
	t.MyType = "TypeConf"
	prefix := t.MyType + ".New"
	t.Initialized = true

	t.setDefault()
	ezlog.Debug().N(prefix).N("Default").Lm(t).Out()

	t.readFileConf()
	ezlog.Debug().N(prefix).N("Raw").Lm(t).Out()

	t.expand()
	ezlog.Debug().N(prefix).N("Expand").Lm(t).Out()
}

func (t *TypeConf) readFileConf() {
	// prefix := t.MyType + ".readFileConf"

	viper.SetConfigType("json")
	viper.SetConfigFile(file.TildeEnvExpand(t.FileConf))
	viper.AutomaticEnv()
	t.Err = viper.ReadInConfig()

	if t.Err == nil {
		t.Err = viper.Unmarshal(&t)
	}
}

func (t *TypeConf) setDefault() {
	if t.FileConf == "" {
		t.FileConf = Default.FileConf
	}
}

func (t *TypeConf) expand() {
	t.FileConf = file.TildeEnvExpand(t.FileConf)
}
