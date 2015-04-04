package disk

import "testing"
import "os"
import "math/rand"
import "bytes"

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

func randomPattern(p [blockSize]byte) {
	for i := 0; i < len(p); i++ {
		p[i] = byte(rand.Int())
	}
}

func TestDiskReadWrite(t *testing.T) {
	var d, err = CreateDisk(filePath, numBlocks, blockSize)
	var p [blockSize]byte
	var p1 [blockSize]byte
	var small [10]byte

	if err != nil {
		t.Fatalf("Failed to Create disk!")
	}

	randomPattern(p)
	d.Write(1, p[:])
	d.Read(1, p1[:])

	if bytes.Compare(p[:], p1[:]) != 0 {
		t.Fatalf("Read Write comparison failed!")
	}

	err = d.Read(numBlocks+1, p[:])
	if err != ErrInvalidBlockNumber {
		t.Fatalf("Expected error for invalid block to read", err)
	}

	err = d.Write(numBlocks+1, p[:])
	if err != ErrInvalidBlockNumber {
		t.Fatalf("Expected error for invalid block to write", err)
	}

	err = d.Read(1, small[:])
	if err != ErrInsufficientBufferSpace {
		t.Fatalf("Expected error, buffer too small", err)
	}

	d.Close()
}
