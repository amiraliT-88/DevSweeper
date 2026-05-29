package scanner

import (
	"os"
	"path/filepath"
	"sync"
)

var Targets = []string{
	"Temp",
	"tmp",
	"Cache",
	"cache",
	".cache",
	"logs",
	"Log",
	"Crashpad",
	"CrashReports",
	"GPUCache",
	"Code Cache",
	"CachedData",
	"__pycache__",
	".pytest_cache",
	".vs",
}

var IgnoredDirs = []string{
	".git",
}

type FoundDir struct {
	Path string
	Size int64
}

func isTarget(name string) bool {
	for _, t := range Targets {
		if name == t {
			return true
		}
	}
	return false
}



func Scan(root string) []FoundDir {
	var results []FoundDir
	var mu sync.Mutex
	var wg sync.WaitGroup

	var walk func(path string)
	walk = func(path string) {
		defer wg.Done()

		entries, err := os.ReadDir(path)
		if err != nil {
			return
		}

		for _, e := range entries {
			if !e.IsDir() {
				continue
			}

			// Ignore .git folders and other safe-zones
			ignored := false
			for _, ig := range IgnoredDirs {
				if e.Name() == ig {
					ignored = true
					break
				}
			}
			if ignored {
				continue
			}

			fullPath := filepath.Join(path, e.Name())

			if isTarget(e.Name()) {
				mu.Lock()
				results = append(results, FoundDir{
					Path: fullPath,
					Size: 0, // Size will be calculated asynchronously
				})
				mu.Unlock()
				continue
			}

			wg.Add(1)
			go walk(fullPath)
		}
	}

	wg.Add(1)
	walk(root)
	wg.Wait()

	return results
}

func DeleteDir(path string) error {
	return os.RemoveAll(path)
}
