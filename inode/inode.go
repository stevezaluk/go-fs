package inode

import (
	"time"
)

const (
	I_FILE = 0
	I_DIR  = 1
)

/*
IndexNode - Represents an Index Node containing metadata about the file.
Associated metadata includes its inode number (a unique identifier) for the index
node, its size (the direct size of its metadata + the data stored in its data blocks),
and timestamps for creation, modification, and access.

In addition to these attributes, a slice of 64-bit integers is provided to serve as a pointers to disk blocks.
We don't use pointers to actual structures in memory here to ensure that we can load an index node independently,
of its blocks. While this is not supremely useful here (as this is all in memory, and the disk blocks need to be loaded
anyway; not saving us any space), this would be useful if this project gets abstracted to use a UNIX file to represent
its disk.

Uniquely, a file's name is not stored in and inode and is only stored by its directory
*/
type IndexNode struct {
	inodeNum  int64
	inodeSize int64
	inodeType int8

	creationTime *time.Time
	modifiedTime *time.Time
	accessTime   *time.Time

	pointers []int64
}

/*
NewIndexNode - Constructor for creating an index node. An inode number
(a unique identifier for the inode) is required here as the constructor
is unaware of the inode table and what inode number should be allocated
next.
*/
func NewIndexNode(inodeNum int64, inodeType int8) *IndexNode {
	timestamp := time.Now()

	return &IndexNode{
		inodeNum:     inodeNum,
		inodeSize:    0,
		inodeType:    inodeType,
		creationTime: &timestamp,
		modifiedTime: &timestamp,
		accessTime:   &timestamp,
	}
}
