package disk

import "os"
import "errors"

var InvalidBuffer = errors.New("Buffer not expected size")

const DefaultBlockSize = 4096

type Disk interface {
	BlockSize() uint32
	NumBlocks() uint64
	Read(b uint64, p []byte) error
	Write(b uint64, p []byte) error
	Close() error
}

type disk struct {
	f         *os.File
	length    uint64
	numBlocks uint64
	blockSize uint32
	path      string
	Disk
}

func OpenDisk(p string, numBlocks uint64) (d *disk, err error) {
	d = new(disk)
	d.path = p
	d.numBlocks = numBlocks
	d.blockSize = DefaultBlockSize
	d.f, err = os.OpenFile(p, os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *disk) NumBlocks() (numBlocks uint64) {
	return d.numBlocks
}

func (d *disk) BlockSize() (blockSize uint32) {
	return d.blockSize
}

func (d *disk) Read(b uint64, p []byte) error {
	return nil
}

func (d *disk) Write(b uint64, p []byte) error {
	return nil
}

func (d *disk) Close() error {
	return nil
}
