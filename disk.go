package main

import (
	"syscall"
)

// Creating structure for DiskStatus
type DiskStatus struct {
	All  uint64 `json:"All"`
	Used uint64 `json:"Used"`
	Free uint64 `json:"Free"`
}

// Function to get
// disk usage of path/disk
func (d *DiskStatus) diskUsage(path string) error {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return err
	}
	d.All = fs.Blocks * uint64(fs.Bsize)
	d.Free = fs.Bfree * uint64(fs.Bsize)
	d.Used = d.All - d.Free
	return nil
}
