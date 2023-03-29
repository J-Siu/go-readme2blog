# go-readme2blog [![Paypal donate](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/donate/?business=HZF49NM9D35SJ&no_recurring=0&currency_code=CAD)

Rapid sync between markdown blog with README.md.

### Table Of Content
<!-- TOC -->

- [Table Of Content](#table-of-content)
- [Features](#features)
- [Doc](#doc)
- [Usage](#usage)
- [Repository](#repository)
- [Contributors](#contributors)
- [Change Log](#change-log)
- [License](#license)

<!-- /TOC -->
<!--more-->
### Features

- Directory mode
  - Sync  
Pair files in a blog(markdown) directory, with repository name in another folder. If the split-marker is found in blog and readme, it combines top(before marker) from blog and bottom(after marker) from readme file, and output to a new location.
  - Check  
Check split and skip markers in files inside folders.
  - List  
Output paring result only.
- File mode
  - Sync  
If the split-marker is found in blog and readme, it combines top(before marker) from blog and bottom(after marker) from readme file, and output to a new location.
  - Check  
Check split and skip markers in files.

### Doc

- https://pkg.go.dev/github.com/J-Siu/go-readme2blog

### Usage

```txt
Sync Blog with README.md

Usage:
  go-sync-readme-blog [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dir         Directory mode
  file        File mode
  help        Help about any command

Flags:
  -d, --debug                 Enable debug output
  -F, --force                 Enable overwriting original file
  -h, --help                  help for go-sync-readme-blog
      --md-ext string         Markdown extension (default ".md")
      --no-error              Do not print error
  -n, --no-parallel           Do not process in parallel
      --no-skip               Ignore skip marker
      --readme string         Readme filename (default "README.md")
  -l, --show-files            Show file lists
      --skip-marker string     (default "<!--skip-sync-->")
      --split-marker string    (default "<!--more-->")

Use "go-sync-readme-blog [command] --help" for more information about a command.
```

Example: Check README.md to blog paring
```sh
go-readme2blog dir list --dir-blog ~/code/public/johnsiu.com/content/blog --dir-src ~/code/public
```

Example: Directory mode sync README.md to blog
```sh
go-readme2blog dir sync --dir-blog ~/code/public/johnsiu.com/content/blog --dir-src ~/code/public --dir-out ~/code/tmp/md
```

### Repository

- [go-readme2blog](https://github.com/J-Siu/go-readme2blog)

### Contributors

- [John Sing Dao Siu](https://github.com/J-Siu)

### Change Log
- v1.0.0
  - Feature complete
- v1.0.1
  - Update to Go 1.20 and dependency
### License

The MIT License

Copyright (c) 2023 John Siu

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.