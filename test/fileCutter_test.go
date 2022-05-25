package test

import (
	"testing"

	"github.com/J-Siu/go-helper"
	"github.com/J-Siu/go-readme2blog/lib"
)

func printCutterInfo(cutter *lib.FileCutter) {
	var lenBottom string = "0"
	var lenTop string = "0"
	if cutter.Bottom != nil {
		lenBottom = *helper.AnyToJsonMarshalSp(len(*cutter.Bottom), false)
	}
	if cutter.Top != nil {
		lenTop = *helper.AnyToJsonMarshalSp(len(*cutter.Top), false)
	}
	helper.Report(cutter, "cutter", false, false)
	helper.Report(cutter.Split, "      Split", false, true)
	helper.Report(lenTop, "   Top line", false, true)
	helper.Report(cutter.SplitLine, " Split line", false, true)
	helper.Report(lenBottom, "Bottom line", false, true)
	helper.Report(helper.Errs, "Errors", true, false)
}

func TestMarkerOnly(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	cutter.Filename = "test-marker-only.md"
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP
	cutter.Read("")
	printCutterInfo(&cutter)
	if !cutter.Split || cutter.SplitLine != 1 {
		t.Fatalf("cutter not correct.")
	}
}
func TestMarkerFirstLn(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	cutter.Filename = "test-marker-first-ln.md"
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP
	cutter.Read("")
	printCutterInfo(&cutter)
	if !cutter.Split || cutter.SplitLine != 1 || len(*cutter.Bottom) != 6 {
		t.Fatalf("cutter not correct.")
	}
}
func TestMarkerLastLn(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	cutter.Filename = "test-marker-last-ln.md"
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP
	cutter.Read("")
	printCutterInfo(&cutter)
	if !cutter.Split || cutter.SplitLine != 7 || len(*cutter.Bottom) != 0 || len(*cutter.Top) != 6 {
		t.Fatalf("cutter not correct.")
	}
}
func TestMarkerNone(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	cutter.Filename = "test-marker-none.md"
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP
	cutter.Read("")
	printCutterInfo(&cutter)
	if cutter.Split || cutter.SplitLine != 0 || cutter.Bottom != nil || cutter.Top != nil {
		t.Fatalf("cutter not correct.")
	}
}
func TestReadBottom(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	cutter.Filename = "test-09.md"
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP
	cutter.ReadBottom("")
	printCutterInfo(&cutter)
	if !cutter.Split || cutter.Bottom == nil || cutter.Top != nil {
		t.Fatalf("cutter not correct.")
	}
}
func TestReadBottomTop(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	var file1 string = "test-09.md"
	var file2 string = "test-az.md"
	cutter.Filename = file1
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP
	cutter.ReadBottom("")
	printCutterInfo(&cutter)
	cutter.ReadTop(file2)
	printCutterInfo(&cutter)
	if !cutter.Split ||
		cutter.Filename != file2 ||
		cutter.Bottom == nil ||
		cutter.Top == nil {
		t.Fatalf("cutter not correct.")
	}
}
func TestReadTop(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	cutter.Filename = "test-az.md"
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP
	cutter.ReadTop("")
	printCutterInfo(&cutter)
	if !cutter.Split || cutter.Bottom != nil || cutter.Top == nil {
		t.Fatalf("cutter not correct.")
	}
}
func TestReadTopBottom(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	var file1 string = "test-09.md"
	var file2 string = "test-az.md"
	cutter.Filename = file1
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP
	cutter.ReadTop("")
	printCutterInfo(&cutter)
	cutter.ReadBottom(file2)
	printCutterInfo(&cutter)
	if !cutter.Split ||
		cutter.Filename != file2 ||
		cutter.Bottom == nil ||
		cutter.Top == nil {
		t.Fatalf("cutter not correct.")
	}
}
func TestSave(t *testing.T) {
	helper.Errs.Clear()
	var cutter lib.FileCutter
	var file1 string = "test-09.md"
	var file2 string = "test-az.md"
	var file3 string = "test-09-az.md"
	var file1TopLine, file2BottomLine, file3TopLine, file3BottomLine int
	cutter.SplitMarker = lib.DEFAULT_MARKER_SPLIT
	cutter.SkipMarker = lib.DEFAULT_MARKER_SKIP

	// Join file1 top, file2 bottom
	cutter.ReadTop(file1)
	file1TopLine = len(*cutter.Top)
	printCutterInfo(&cutter)
	cutter.ReadBottom(file2)
	file2BottomLine = len(*cutter.Bottom)
	printCutterInfo(&cutter)
	// Save to file3
	cutter.Save(file3)
	printCutterInfo(&cutter)
	// Read file3
	cutter.Read(file3)
	file3BottomLine = len(*cutter.Bottom)
	file3TopLine = len(*cutter.Top)
	printCutterInfo(&cutter)
	if !cutter.Split ||
		cutter.Filename != file3 ||
		file3BottomLine != file2BottomLine ||
		file3TopLine != file1TopLine ||
		cutter.Bottom == nil ||
		cutter.Top == nil {
		t.Fatalf("cutter not correct.")
	}
}
