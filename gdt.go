package gdt

import (
	"unsafe"
	"video"
)

type GDTDesc struct {
	Base  uint32
	Limit uint32
	Type  uint8
}

//var GDT uint64
const size uint16 = 5

var Table [size][8]uint8

var tss [27]uint32

var err [40]byte

func Pack(desc GDTDesc) (g [8]uint8) {
	if desc.Limit > 65536 && (desc.Limit&0xFFF != 0xFFF) {
		video.Error(err, int(desc.Limit), true)
	}
	if desc.Limit > 65536 {
		desc.Limit = desc.Limit >> 12
		g[6] = 0xC0
	} else {
		g[6] = 0x40
	}

	g[0] = uint8(desc.Limit)
	g[1] = uint8(desc.Limit >> 8)
	g[6] |= uint8(desc.Limit >> 16)

	g[2] = uint8(desc.Base)
	g[3] = uint8(desc.Base >> 8)
	g[4] = uint8(desc.Base >> 16)
	g[7] = uint8(desc.Base >> 24)

	g[5] = desc.Type
	return
}

func SetupGDT() {
	for i := 0; i < len("GDT entry too large"); i++ {
		err[i] = "GDT entry too large"[i]
	}

	loadTable()
	loadGDT(&Table, unsafe.Sizeof(Table))
	reloadSegments()
}

//extern __load_gdt
func loadGDT(*[size][8]uint8, uintptr)

//extern __reload_segments
func reloadSegments()

func loadTable() {
	Table[0] = Pack(GDTDesc{Base: 0, Limit: 0, Type: 0})
	Table[1] = Pack(GDTDesc{Base: 0, Limit: 0xFFFFFFFF, Type: 0x9A})
	Table[2] = Pack(GDTDesc{Base: 0, Limit: 0xFFFFFFFF, Type: 0x92})
	Table[3] = Pack(GDTDesc{Base: uint32(uintptr(unsafe.Pointer(&tss[0]))), Limit: uint32(unsafe.Sizeof(tss)), Type: 0x89})
}
