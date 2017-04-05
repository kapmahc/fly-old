package web

import (
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/google/uuid"
)

// Uploader attachment uploader
type Uploader interface {
	Save(multipart.File, *multipart.FileHeader) (string, int64, error)
	Remove(string) error
}

// NewFileSystemUploader new file-system uploader
func NewFileSystemUploader(root, home string) (Uploader, error) {
	err := os.MkdirAll(root, 0755)
	return &FileSystemUploader{home: home, root: root}, err
}

// FileSystemUploader file-system storage
type FileSystemUploader struct {
	home string
	root string
}

// Remove remove file
func (p *FileSystemUploader) Remove(url string) error {
	return os.Remove(path.Join(p.root, url[len(p.home)+1:]))
}

// Save save file to file-system
func (p *FileSystemUploader) Save(fd multipart.File, fh *multipart.FileHeader) (string, int64, error) {
	name := uuid.New().String() + path.Ext(fh.Filename)
	dst, err := os.Create(path.Join(p.root, name))
	if err != nil {
		return "", 0, err
	}
	defer dst.Close()
	size, err := io.Copy(dst, fd)
	return p.home + "/" + name, size, err
}
