package partition

// Partition is a struct representing a partition on a disk
type Partition struct {
	Prev *Partition
	New  *Partition
}

// New returns a pointer to a new Partition
func New(t Type, fsType *SysType, start, end uint) *Partition {
	return nil
}
