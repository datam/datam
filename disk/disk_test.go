package disk

import "testing"

func TestDiskCreate(t *testing.T) {
	var d, err = OpenDisk("/tmp/disk.file", 1000)

	if d == nil || err != nil {
		t.Fatalf("Failed to create disk!")
	}
}
