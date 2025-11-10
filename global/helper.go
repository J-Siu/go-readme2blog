package global

import (
	"path"

	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/J-Siu/go-helper/v2/str"
	"github.com/J-Siu/go-readme2blog/lib"
)

// Helper functions that are Flag dependent

// Check file contain skip marker and split marker
func ChkMarker(filePath string) (hasSkip, hasSplit bool, e error) {
	var (
		content *[]string
	)
	content, e = file.ReadStrArray(filePath)
	if e == nil {
		hasSkip = str.ArrayContains(content, &Conf.MarkerSkip, false)
		hasSplit = str.ArrayContains(content, &Conf.MarkerSplit, false)
	}
	return hasSkip, hasSplit, e
}

// Map readmes in all dirs in `dirRepo` to blog in `dirBlog`
func MapReadmeBlog(dirBlog, dirRepo string) *lib.FileMap {
	MapBlog := make(lib.FileMap)
	MapBlog.MapFileDir(dirBlog)
	MapReadme := make(lib.FileMap)
	MapReadme.MapDirFile(dirRepo, DEFAULT_README)
	return MapReadme.Join(&MapBlog)
}

func MapReadmeBlogSync(mapReadmeBlog *lib.FileMap, dirOut string) {
	if mapReadmeBlog != nil {
		var (
			fileStich = new(lib.FileStitch)
			property  = new(lib.FileStitchProperty)
		)
		property.MarkerSkip = &Conf.MarkerSkip
		property.MarkerSplit = &Conf.MarkerSplit

		for readme, blog := range *mapReadmeBlog {
			fileOut := path.Join(dirOut, path.Base(blog))
			property.FileBottom = &readme
			property.FileOut = &fileOut
			property.FileTop = &blog
			fileStich.New(property).Run()
			ezlog.Log().N("Processed").M(readme).M("->").M(blog).Out()
			fileStich.Reset()
		}
	}
}
