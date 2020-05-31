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

type TSS32 struct {
	// _        uint32
	Backlink uint32
	Esp0     uint32
	Ss0      uint32
	Esp1     uint32
	Ss1      uint32
	Esp2     uint32
	Ss2      uint32
	Cr3      uint32
	Eip      uint32
	Eflags   uint32
	Eax      uint32
	Ecx      uint32
	Edx      uint32
	Ebx      uint32
	Esp      uint32
	Ebp      uint32
	Esi      uint32
	Edi      uint32
	Es       uint32
	Cs       uint32
	Ss       uint32
	Ds       uint32
	Fs       uint32
	Gs       uint32
	Ldtr     uint32
	Iomap    uint32
	// _        uint32
}

//var GDT uint64
//the number of GDT entry
const size uint16 = 5
//every GDT entry is 8 byte
var Table [size][8]uint8

var TSSA TSS32
var TSSB TSS32

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
	// loadTR(3 * 8)
}

//extern __load_tr
func loadTR(tr int)

func LoadTR(tr int) {
	loadTR(tr)
}

//extern __load_gdt
func loadGDT(*[size][8]uint8, uintptr)

//extern __reload_segments
func reloadSegments()

func loadTable() {
	TSSA.Ldtr = 0
	TSSA.Iomap = 0x40000000

	TSSB.Ldtr = 0
	TSSB.Iomap = 0x40000000

	Table[0] = Pack(GDTDesc{Base: 0, Limit: 0, Type: 0})
	Table[1] = Pack(GDTDesc{Base: 0, Limit: 0xFFFFFFFF, Type: 0x9A}) // code
	Table[2] = Pack(GDTDesc{Base: 0, Limit: 0xFFFFFFFF, Type: 0x92}) // data
	Table[3] = Pack(GDTDesc{Base: uint32(uintptr(unsafe.Pointer(&TSSA.Backlink))), Limit: uint32(unsafe.Sizeof(TSSA)), Type: 0x89})
	Table[4] = Pack(GDTDesc{Base: uint32(uintptr(unsafe.Pointer(&TSSB.Backlink))), Limit: uint32(unsafe.Sizeof(TSSB)), Type: 0x89})

}

//extern __stack_ptr
func stack() uint32

func Stack() uint32 {
	stackptr := stack()
	video.PrintHex(uint64(stackptr), false, true, true, 8)
	return stackptr
}
