package go_file

import (
	"os"
	"path"
)

type File struct {
	Path     string
	FileName string
	Content  Content
}

func (f File) Create() {
	_, err := os.Stat(f.Path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(f.Path, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	fileName := path.Join(f.Path, f.FileName)

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(f.Content.ToString())
	if err != nil {
		panic(err)
	}
}
