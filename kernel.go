package kernel

import (
	//	"color"
	"idt"
	"unsafe"
	"video"

	//"ptr"
	"gdt"
	"pit"

	//"unsafe"
	"kbd"
)

//extern __test_int
func testInt()

//extern __test_args
func testArgs(c rune)

var MULTIBOOT_BOOTLOADER_MAGIC uint32 = 0x1BADB002

//func Kmain() {
func Kmain(mdb uintptr, magic uintptr) {
	video.Init()
	video.Clear()
	video.PrintHex(uint64(*(*uint32)(unsafe.Pointer(magic))), false, true, true, 8)
	// if magic != MULTIBOOT_BOOTLOADER_MAGIC {
	// 	video.Print("Invalid magic number\n")
	// }
	gdt.SetupGDT()
	idt.SetupIDT()
	idt.SetupIRQ()
	pit.Init()
	kbd.Init()
	video.Print("Hello kernel\n")

	for i := 0; i < 16; i++ {
		video.Print("abc\n")
	}
}
