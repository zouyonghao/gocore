package kernel

import (
	"gdt"
	"idt"
	"kbd"
	"pit"
	"ptr"
	"video"
	"regs"
	"unsafe"
)

//extern __test_int
func testInt()

//extern __test_args
func testArgs(c rune)

// var MULTIBOOT_BOOTLOADER_MAGIC uint32 = 0x1BADB002

//extern __taskswitch3
func taskswitch3()

//extern __taskswitch4
func taskswitch4()

//extern __syscall
func syscall()

func syscall_imp(r *regs.Regs) {
    video.Print("syscall\n")
}

func taskB() {
	for i := 0; i < 10000; i++ {
	    syscall()
		video.Print("This is Task B\n")
	}
	taskswitch3()
}

//extern sys_malloc
func sysAlloc(n uintptr) uint32

//extern __kernel_end
func kernelEnd() uintptr

func Kmain() {
	// func Kmain(mdb uintptr, magic uintptr) {
	video.Init()
	video.Clear()
    video.Print(`    mmmm     mmmm       mmmm     mmmm    mmmmmm    mmmmmmmm` + "\n")
    video.Print(`  ##""""#   ##""##    ##""""#   ##""##   ##""""##  ##""""""` + "\n") 
    video.Print(` ##        ##    ##  ##"       ##    ##  ##    ##  ##` + "\n")
    video.Print(` ##  mmmm  ##    ##  ##        ##    ##  #######   #######` + "\n")
    video.Print(` ##  ""##  ##    ##  ##m       ##    ##  ##  "##m  ##` + "\n")
    video.Print(`  ##mmm##   ##mm##    ##mmmm#   ##mm##   ##    ##  ##mmmmmm ` + "\n")
    video.Print(`    """"     """"       """"     """"    ""    """ """"""""` + "\n")
	// video.PrintHex(uint64(*(*uint32)(unsafe.Pointer(magic))), false, true, true, 8)
	// if magic != MULTIBOOT_BOOTLOADER_MAGIC {
	// 	video.Print("Invalid magic number\n")
	// }

	gdt.SetupGDT()
	idt.SetupIDT()
	idt.SetupIRQ()
	// page.Init()
	pit.Init()
	kbd.Init()
	video.Print("Hello kernel\n")
	
	dummy := syscall_imp
	idt.AddIRQ(15, **(**uintptr)(unsafe.Pointer(&dummy)))
	
	gdt.LoadTR(3 * 8)

	gdt.TSSB.Eip = uint32(ptr.FuncToPtr(taskB))
	gdt.TSSB.Eflags = 0x00000202 | (1 << 9)
	gdt.TSSB.Eax = 0
	gdt.TSSB.Ecx = 0
	gdt.TSSB.Edx = 0
	gdt.TSSB.Ebx = 0
	gdt.TSSB.Esp0 = uint32(kernelEnd() + 4096)
	gdt.TSSB.Esp = uint32(kernelEnd() + 4096)
	// gdt.TSSB.Esp = sysAlloc(10)
	gdt.TSSB.Ebp = 0
	gdt.TSSB.Esi = 0
	gdt.TSSB.Edi = 0

	gdt.TSSB.Ss0 = 2 * 8
	gdt.TSSB.Cs = 1 * 8

	gdt.TSSB.Es = 2 * 8
	gdt.TSSB.Ss = 2 * 8
	gdt.TSSB.Ds = 2 * 8
	gdt.TSSB.Fs = 2 * 8
	gdt.TSSB.Gs = 2 * 8

	var tmp int = 0
	for {
		video.Print("This is Task A\n")
		if tmp%1000000 == 999 {
			taskswitch4()
		}
		tmp++
	}
	// taskswitch3()

}
