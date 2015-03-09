package disk

import "os"
import "errors"

var errInvalidBuffer = errors.New("Buffer not expected size")
var errInvalidDiskSize = errors.New("Disk length aligned to block size")
var errInvalidBlockNumber = errors.New("Block Number is beyond max length")
var errInsufficientBufferSpace = errors.New("Buffer provided is not the expected length")

const DefaultBlockSize = 4096

type Disk interface {
	BlockSize() uint64
	NumBlocks() uint64
	Read(b uint64, p []byte) error
	Write(b uint64, p []byte) error
	Close() error
}

type disk struct {
	f         *os.File
	length    uint64
	numBlocks uint64
	blockSize uint64
	path      string
	Disk
}

func CreateDisk(p string, numBlocks uint64, blockSize uint64) (d *disk, err error) {
	f, err := os.Create(p)
	if err != nil {
		f.Close()
		return nil, err
	}

	err = f.Truncate(int64(numBlocks * blockSize))
	if err != nil {
		f.Close()
		return nil, err
	}

	f.Close()

	return OpenDisk(p)
}

func OpenDisk(p string) (d *disk, err error) {
	d = new(disk)
	d.path = p
	d.blockSize = DefaultBlockSize

	d.f, err = os.OpenFile(p, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	fileInfo, err := d.f.Stat()
	if err != nil {
		d.f.Close()
		return nil, err
	}

	d.length = uint64(fileInfo.Size())
	if err != nil {
		return nil, err
	}

	if d.length%d.blockSize != 0 {
		return nil, errInvalidDiskSize
	}

	d.numBlocks = d.length / d.blockSize
	return d, nil
}

func (d *disk) NumBlocks() (numBlocks uint64) {
	return d.numBlocks
}

func (d *disk) BlockSize() (blockSize uint64) {
	return d.blockSize
}

func (d *disk) Read(b uint64, p []byte) error {
	if b > d.numBlocks {
		return errInvalidBlockNumber
	}

	if len(p) < int(d.blockSize) {
		return errInsufficientBufferSpace
	}

	_, err := d.f.ReadAt(p, int64(d.blockSize*b))
	if err != nil {
		return err
	}

	return nil
}

func (d *disk) Write(b uint64, p []byte) error {
	if b > d.numBlocks {
		return errInvalidBlockNumber
	}

	if len(p) != int(d.blockSize) {
		return errInsufficientBufferSpace
	}

	_, err := d.f.WriteAt(p, int64(b))
	if err != nil {
		return err
	}

	return nil
}

func (d *disk) Close() error {
	return d.f.Close()
}
