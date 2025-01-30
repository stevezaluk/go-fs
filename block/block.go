package block

/*
DataBlock - Represents a block of data that an index node might point to. Its
data is stored as a byte array to ensure that this representation of one is as
close to an actual filesystem as possible

In addition to the byte array that is provided in this struct, an array of
DataBlock pointers is also provided for implementing indirect block pointers
(A pointer to a block of pointers)
*/
type DataBlock struct {
	blockNum int64
	data     []byte
	pointers []*DataBlock
}

/*
NewDataBlock - A constructor for the data block. The 'pointers' parameter is provided here to initialize
the block with additional block pointers. If this is nil, then a byte array with a maximum capacity of 4096
is created. This is only done when 'pointers' is nil as a data block cannot hold both data for an inode
and pointers to indirect blocks.
*/
func NewDataBlock(blockNum int64, pointers []*DataBlock) *DataBlock {
	ret := &DataBlock{blockNum: blockNum}

	if pointers == nil {
		ret.data = make([]byte, 0, 4096)
	}

	return ret
}
