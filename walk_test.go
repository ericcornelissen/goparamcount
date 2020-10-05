package main

import (
	"path/filepath"
	"testing"
)

func TestSkipDir(t *testing.T) {
	t.Run("recursive disabled", func(t *testing.T) {
		result := skipDir("./directory", false)
		if result != filepath.SkipDir {
			t.Error("Expected result to be `filepath.SkipDir`")
		}
	})
	t.Run("random directory (recursive)", func(t *testing.T) {
		result := skipDir("./directory", true)
		if result != nil {
			t.Error("Expected result to be `nil`")
		}
	})
	t.Run("directories always excluded (recursive)", func(t *testing.T) {
		result := skipDir("./.git", true)
		if result != filepath.SkipDir {
			t.Error("Expected result to be `filepath.SkipDir`")
		}

		result = skipDir("./testdata", true)
		if result != filepath.SkipDir {
			t.Error("Expected result to be `filepath.SkipDir`")
		}

		result = skipDir("./vendor", true)
		if result != filepath.SkipDir {
			t.Error("Expected result to be `filepath.SkipDir`")
		}
	})
}

func TestSkipFile(t *testing.T) {
	var noExcludePatterns []string

	t.Run("non-.go file", func(t *testing.T) {
		result := skipFile("file.txt", noExcludePatterns)
		if result == false {
			t.Error("Expected result to be true")
		}
	})
	t.Run(".go file", func(t *testing.T) {
		result := skipFile("file.go", noExcludePatterns)
		if result == true {
			t.Error("Expected result to be false")
		}
	})
	t.Run("_test.go file", func(t *testing.T) {
		result := skipFile("file_test.go", noExcludePatterns)
		if result == false {
			t.Error("Expected result to be true")
		}
	})
	t.Run("custom exclude pattern", func(t *testing.T) {
		excludePatterns := []string{"foo*.go"}
		result := skipFile("foobar.go", excludePatterns)
		if result == false {
			t.Error("Expected result to be true")
		}
	})
}
