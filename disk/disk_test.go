package disk

import "testing"
import "os"

const numBlocks = 1000
const blockSize = 4096
const filePath = "/tmp/disk.file"

func TestDiskCreate(t *testing.T) {
	var d, err = CreateDisk(filePath, numBlocks, blockSize)

	if d == nil || err != nil {
		t.Fatalf("Failed to create disk!")
	}

	d.Close()

	d, err = OpenDisk(filePath)

	if d == nil || err != nil {
		t.Fatalf("Failed to open disk!")
	}

	if d.NumBlocks() != numBlocks {
		t.Fatalf("numBlocks mismatch!")
	}

	d.Close()

	err = os.Remove("/tmp/disk.file")
	if err != nil {
		t.Fatalf("Failed to remove disk file after test")
	}
}

func TestDiskReadWrite(t *testing.T) {

}
