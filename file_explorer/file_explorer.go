package file_explorer

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type FileExplorer struct {
	CurrentPath string
	DirEntries []fs.DirEntry
}

func NewFileExplorer() (FileExplorer, error){
	f := FileExplorer{}
	currentPath, err := os.Getwd()
	if err != nil {
		return f, err
	}
	f.CurrentPath = currentPath
	dirEntries, err := os.ReadDir(currentPath)
	if err != nil {
		return f, err
	}
	f.DirEntries = dirEntries
	return f, nil
}

func (fe *FileExplorer) GetCurrentPath() string{
	return fe.CurrentPath
}

func (fe *FileExplorer) GetDirectoryEntries() []fs.DirEntry{
	return fe.DirEntries
}

func (fe *FileExplorer) SearchInPath(keyword string) []fs.DirEntry{
	var ret []fs.DirEntry
	for _, entry := range fe.DirEntries {
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
	err = os.Chdir(absolutePath)
	if err != nil {
		return err
	}
	fe.CurrentPath = absolutePath
	dirEntries, err := os.ReadDir(absolutePath)
	if err != nil {
		return err
	}
	fe.DirEntries = dirEntries
	return nil
}