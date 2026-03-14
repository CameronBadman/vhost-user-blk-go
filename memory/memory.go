package memory

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Region struct {
	GuestPhysAddr uint64
	Size          uint64
	UserspaceAddr uint64
	MmapOffset    uint64
	HostAddr      uintptr
}

type Table struct {
	Regions []Region
}

// TODO ATOMICITY IS NEEDED HERE
func ParseAndMap(payload []byte, fds []int) (*Table, error) {
	nregions := binary.LittleEndian.Uint32(payload[0:4])
	regions := make([]Region, nregions)
	for i := 0; i < int(nregions); i++ {
		off := 8 + i*32
		r := Region{
			GuestPhysAddr: binary.LittleEndian.Uint64(payload[off+0:]),
			Size:          binary.LittleEndian.Uint64(payload[off+8:]),
			UserspaceAddr: binary.LittleEndian.Uint64(payload[off+16:]),
			MmapOffset:    binary.LittleEndian.Uint64(payload[off+24:]),
		}
		addr, _, errno := unix.Syscall6(
			unix.SYS_MMAP,
			uintptr(r.UserspaceAddr),
			uintptr(r.Size+r.MmapOffset),
			unix.PROT_READ|unix.PROT_WRITE,
			unix.MAP_SHARED,
			uintptr(fds[i]),
			0,
		)
		if errno != 0 {
			return nil, fmt.Errorf("mmap region %d: %w", i, errno)
		}
		r.HostAddr = addr
		unix.Close(fds[i])
		regions[i] = r
	}
	return &Table{Regions: regions}, nil
}

func (t *Table) GuestToHost(gpa uint64) (uintptr, bool) {
	for i := range t.Regions {
		r := &t.Regions[i]
		if gpa >= r.GuestPhysAddr && gpa < r.GuestPhysAddr+r.Size {
			return r.HostAddr + uintptr(r.MmapOffset) + uintptr(gpa-r.GuestPhysAddr), true
		}
	}
	return 0, false
}

func (t *Table) Unmap() {
	for i := range t.Regions {
		r := &t.Regions[i]
		unix.Munmap((*[1 << 30]byte)(unsafe.Pointer(r.HostAddr))[:r.Size+r.MmapOffset])
	}
	t.Regions = nil
}
