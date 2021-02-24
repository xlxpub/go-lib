package filex

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ErrorMaxModCnt = errors.New("max modify cnt reached")

func RenameSuffix(count int, dir, oldSuffix, newSuffix string) (int, error) {
	modCnt := 0
	maxModCnt := count
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, oldSuffix) {
			parts := strings.Split(path, ".")
			suffix := parts[len(parts)-1]
			suffix = strings.Replace(suffix, oldSuffix, newSuffix, 1)

			parts[len(parts)-1] = suffix
			newPath := strings.Join(parts, ".")
			if newPath != path {
				log.Printf("DEBUG renaming %s to %s\n", path, newPath)
				if err := os.Rename(path, newPath); err != nil {
					return err
				}
				modCnt++
				if modCnt >= maxModCnt {
					return ErrorMaxModCnt
				}
			}

		}
		return nil
	})
	if err != nil {
		return modCnt, err
	}
	return modCnt, nil
}
