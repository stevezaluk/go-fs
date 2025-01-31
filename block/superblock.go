package block

import "github.com/stevezaluk/go-fs/inode"

type SuperBlock struct {
	// fsType - The file system type that has been used. This will always be go-fs
	fsType string

	// totalBlocks - The number of data blocks that can be used
	totalBlocks int64

	// totalInodes - The number of inodes that can be used
	totalInodes int64

	// totalUsedBlocks - The number of data blocks currently in use
	totalUsedBlocks int64

	// totalUsedInodes - The number of inodes currently in use
	totalUsedInodes int64

	// inodeTable - A map that tracks what index nodes are currently in-use
	inodeTable map[int64]*inode.IndexNode

	// freeBlockTable - A map that tracks what blocks are currently in-use
	freeBlockTable map[int64]*DataBlock
}
