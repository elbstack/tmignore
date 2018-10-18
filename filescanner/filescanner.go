package filescanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobwas/glob"
	log "github.com/sirupsen/logrus"
)

const separator = '/'

func FilePathWalkDir(root, pattern string) (map[string]struct{}, error) {
	log.WithFields(log.Fields{"root": root, "pattern": pattern}).Debug("Filepath walker args")
	var g glob.Glob
	g = glob.MustCompile(pattern, separator)

	patternParts := strings.Split(pattern, "/")
	dirName := patternParts[len(patternParts)-1]
	log.WithField("path separator", dirName).Debug("name to split at")

	unique := make(map[string]struct{})

	dirsScanned := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirsScanned++
			fmt.Printf("\033[2K\r%v: %d", "Directories scanned", dirsScanned)
			if g.Match(path) {
				fullPath, _ := filepath.Abs(path)
				fullPathSplitted := strings.SplitAfterN(fullPath, dirName, 2)[0]
				if _, ok := unique[fullPathSplitted]; !ok {
					unique[fullPathSplitted] = struct{}{}
				}
			}
		}

		return nil
	})
	return unique, err
}
