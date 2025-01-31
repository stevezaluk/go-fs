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
	inodeNum int64
	contents map[string]int64
}

/*
NewDirectoryEntry - A constructor for the DirectoryEntry struct. It accepts two arguments
for its parameters: inodeNum should be the inode number that represents this directory, parent
should be the inode number of its previous directory. If -1 is passed to this parameter then
it marks the directory with no parent
*/
func NewDirectoryEntry(inodeNum int64, parent int64) *DirectoryEntry {

	/*
		In the case that there is no parent, then we set our previous inode number
		to our self to ensure that the user can still use the '..' path in operations
		requiring it
	*/
	prevInodeNum := inodeNum
	if parent != -1 {
		prevInodeNum = parent
	}

	/*
		Finally we define our contents map which represents the contents of oru directory.
		We do this so that we can append both '.' and '..' to the map. This ensures ease of
		use when interacting with both the current directory and the parent directory
	*/
	contents := map[string]int64{
		".":  inodeNum,
		"..": prevInodeNum,
	}

	return &DirectoryEntry{inodeNum: inodeNum, contents: contents}
}
