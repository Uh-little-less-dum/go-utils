package file_handlers_base

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
)

type BaseFileHandler struct {
	Path string
}

func (b BaseFileHandler) Exists() bool {
	_, err := os.Stat(b.Path)
	return !errors.Is(err, os.ErrNotExist)
}

// func (b BaseFileHandler) IsDir() bool {
// 	v, err := os.Stat(b.Path)
// 	if err != nil {
// 		return false
// 	}
// 	return v.IsDir()
// }

func (b BaseFileHandler) Read() []byte {
	d, err := os.ReadFile(b.Path)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

// func (b BaseFileHandler) WriteContent(newContent string) {
// 	err := os.WriteFile(b.Path, []byte(newContent), 0777)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func (b BaseFileHandler) WriteContentBytes(newContent []byte) {
	err := os.WriteFile(b.Path, newContent, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// func (a BaseFileHandler) CopyTo(b BaseFileHandler) bool {
// 	if (a.Exists()) && b.Exists() {
// 		inputContent := a.Read()
// 		b.WriteContentBytes(inputContent)
// 		return true
// 	}
// 	return false
// }

// func (a BaseFileHandler) Delete() {
// 	err := os.Remove(a.Path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (a BaseFileHandler) RelativePathTo(b BaseFileHandler) string {
// 	s, err := filepath.Rel(a.Path, b.Path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return s
// }
