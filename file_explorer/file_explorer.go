package file_explorer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type FileExplorer struct {
	currentPath string
	dirEntries []fs.DirEntry
}

func NewFileExplorer() (FileExplorer, error){
	f := FileExplorer{}
	currentPath, err := os.Getwd()
	if err != nil {
		return f, err
	}
	f.currentPath = currentPath
	dirEntries, err := os.ReadDir(currentPath)
	if err != nil {
		return f, err
	}
	f.dirEntries = dirEntries
	return f, nil
}

func (fe FileExplorer) GetCurrentPath() string{
	return fe.currentPath
}

func (fe FileExplorer) GetDirectoryEntries() []fs.DirEntry{
	return fe.dirEntries
}

func (fe FileExplorer) SearchInPath(keyword string) []fs.DirEntry{
	var ret []fs.DirEntry
	for _, entry := range fe.dirEntries {
		if strings.Contains(entry.Name(), keyword) {
			ret = append(ret, entry)
		}
	}
	return ret
}

func (fe *FileExplorer) ChangeDirectory(dir string) error{
	absolutePath, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	fe.currentPath = absolutePath
	dirEntries, err := os.ReadDir(absolutePath)
	if err != nil {
		return err
	}
	fmt.Println("abs path:", absolutePath)
	fe.dirEntries = dirEntries
	return nil
}