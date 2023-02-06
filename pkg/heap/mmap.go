package heap

import (
	"github.com/efficientgo/core/merrors"
	"golang.org/x/sys/unix"
	"os"
)

type MemoryMap struct {
	f *os.File // nil if anonymous.
	b []byte
}

func OpenFileBacked(path string, size int) (mf *MemoryMap, _ error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b, err := unix.Mmap(int(f.Fd()), 0, size, unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		return nil, merrors.New(f.Close(), err).Err()
	}

	return &MemoryMap{f: f, b: b}, nil
}

func (f *MemoryMap) Close() error {
	errs := merrors.New()
	errs.Add(unix.Munmap(f.b))
	errs.Add(f.f.Close())
	return errs.Err()
}

func (f *MemoryMap) Bytes() []byte { return f.b }
