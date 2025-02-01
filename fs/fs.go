package fs

import "github.com/stevezaluk/go-fs/block"

/*
FileSystem - A representation of the file system as a whole. Contains
traditional system calls that you would expect to find in a filesystem
*/
type FileSystem struct {
	SuperBlock *block.SuperBlock
}

/*
NewFileSystem - A constructor for the file system object
*/
func NewFileSystem() *FileSystem {
	return &FileSystem{SuperBlock: block.NewSuperBlock()}
}
