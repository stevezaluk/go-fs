package inode

import (
	"time"
)

/*
IndexNode - Represents an Index Node containing metadata about the file.
Associated metadata includes its inode number (a unique identifier) for the index
node, its size (the direct size of its metadata + the data stored in its data blocks),
and timestamps for creation, modification, and access.

Uniquely, a file's name is not stored in and inode and is only stored by its directory
*/
type IndexNode struct {
	inodeNum  int64
	inodeSize int64

	creationTime *time.Time
	modifiedTime *time.Time
	accessTime   *time.Time

	// data blocks here
}

/*
NewIndexNode - Constructor for creating an index node. An inode number
(a unique identifier for the inode) is required here as the constructor
is unaware of the inode table and what inode number should be allocated
next.
*/
func NewIndexNode(inodeNum int64) *IndexNode {
	timestamp := time.Now()

	return &IndexNode{
		inodeNum:     inodeNum,
		inodeSize:    0,
		creationTime: &timestamp,
		modifiedTime: &timestamp,
		accessTime:   &timestamp,
	}
}
