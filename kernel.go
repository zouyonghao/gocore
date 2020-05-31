package kernel

import (
	//	"color"
	"idt"
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

//func Kmain() {
func Kmain(mdb uintptr, magic uint16) {
	video.Init()
	video.Clear()
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
