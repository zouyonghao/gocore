package pit

import (
	"idt"
	"regs"
	"unsafe"
	"video"
)

var ticks int

var times int

//extern __taskswitch3
func taskswitch3()

//extern __taskswitch4
func taskswitch4()

func handler(r *regs.Regs) {
	ticks++
	if ticks%36 == 0 {
		video.Println("Tick")
	} else if ticks%36 == 18 {
		video.Println("Tock")
		times++
		if times == 1 {
			// taskswitch3()
		}
	}
}

func Init() {
	dummy := handler
	idt.AddIRQ(0, **(**uintptr)(unsafe.Pointer(&dummy)))
	video.Println("PIT Initialized")
}
