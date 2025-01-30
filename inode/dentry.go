package inode

/*
DirectoryEntry - Represents a directory entry. A directory
entry is effectively a list of inode numbers (a unique identifier
for the inode) and its associated name.

Directory entries can be represented as linked lists here (and usually
are in traditional operating systems), however I am using a hash map here
as it provides us better random access (randomly accessing an element in a
hash map is O(1) while it is O(n) in a linked list [in its worst case])
*/
type DirectoryEntry struct {
	inode    *IndexNode
	contents map[string]int64
}

/*
NewDirectoryEntry - A constructor for the DirectoryEntry struct. It accepts a parent
directory for its parameter. If nil is passed into this pointer, then the directory
entry is treated as the root directory
*/
func NewDirectoryEntry(inodeNum int64, parent *DirectoryEntry) *DirectoryEntry {

	/*
		In the case that there is no parent, then we set our previous inode number
		to our self to ensure that the user can still use the '..' path in operations
		requiring it
	*/
	prevInodeNum := inodeNum
	if parent != nil {
		prevInodeNum = parent.inode.inodeNum
	}

	/*
		Once again we need a free inode number returned by the inode table as this structures
		constructor is unaware of what index numbers are free
	*/
	inode := NewIndexNode(inodeNum, I_DIR)

	/*
		Finally we define our contents map which represents the contents of oru directory.
		We do this so that we can append both '.' and '..' to the map. This ensures ease of
		use when interacting with both the current directory and the parent directory
	*/
	contents := map[string]int64{
		".":  inode.inodeNum,
		"..": prevInodeNum,
	}

	return &DirectoryEntry{inode: inode, contents: contents}
}
