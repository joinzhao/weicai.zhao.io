package create

import (
	"os"
	"path/filepath"
)

const (
	CmdName = "create"
)

func createFile(f string) error {
	_, err := os.Lstat(f)
	if err != nil {
		// file not exists
		dir := filepath.Dir(f)
		_, err = os.Lstat(dir)
		if err != nil {
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				return err
			}
		}
		_, err = os.Create(f)

		return err
	}
	return nil
}
