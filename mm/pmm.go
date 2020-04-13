package kernel

// type taskstate struct {
// 	ts_link      uint32  // old ts selector
// 	ts_esp0      uintptr // stack pointers and segment selectors
// 	ts_ss0       uint16  // after an increase in privilege level
// 	ts_padding1  uint16
// 	ts_esp1      uintptr
// 	ts_ss1       uint16
// 	ts_padding2  uint16
// 	ts_esp2      uintptr
// 	ts_ss2       uint16
// 	ts_padding3  uint16
// 	ts_cr3       uintptr // page directory base
// 	ts_eip       uintptr // saved state from last task switch
// 	ts_eflags    uint32
// 	ts_eax       uint32 // more saved state (registers)
// 	ts_ecx       uint32
// 	ts_edx       uint32
// 	ts_ebx       uint32
// 	ts_esp       uintptr
// 	ts_ebp       uintptr
// 	ts_esi       uint32
// 	ts_edi       uint32
// 	ts_es        uint16 // even more saved state (segment selectors)
// 	ts_padding4  uint16
// 	ts_cs        uint16
// 	ts_padding5  uint16
// 	ts_ss        uint16
// 	ts_padding6  uint16
// 	ts_ds        uint16
// 	ts_padding7  uint16
// 	ts_fs        uint16
// 	ts_padding8  uint16
// 	ts_gs        uint16
// 	ts_padding9  uint16
// 	ts_ldt       uint16
// 	ts_padding10 uint16
// 	ts_t         uint16 // trap on task switch
// 	ts_iomb      uint16 // i/o map base address
// }

func gdtInit() {

}

func pmmInit() {
	gdtInit()
}
