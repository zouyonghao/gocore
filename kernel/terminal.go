package kernel

import "unsafe"

const (
	COLOR_BLACK = iota
	COLOR_BLUE
	COLOR_GREEN
	COLOR_CYAN
	COLOR_RED
	COLOR_MAGENTA
	COLOR_BROWN
	COLOR_LIGHT_GREY
	COLOR_DARK_GREY
	COLOR_LIGHT_BLUE
	COLOR_LIGHT_GREEN
	COLOR_LIGHT_CYAN
	COLOR_LIGHT_RED
	COLOR_LIGHT_MAGENTA
	COLOR_LIGHT_BROWN
	COLOR_WHITE
)

const (
	VGA_WIDTH  = 80
	VGA_HEIGHT = 25
)

var row, column, color uint8
var buffer uintptr

func makeColor(fg uint8, bg uint8) uint8 {
	return fg | bg<<4
}

func makeVGAEntry(c byte, color uint8) uint16 {
	return uint16(c) | uint16(color)<<8
}

func terminalInit() {
	row = 1
	column = 0
	color = makeColor(COLOR_WHITE, COLOR_BLACK)
	buffer = 0xB8000
}

func terminalPutEntryAt(c byte, color uint8, column uint8, row uint8) {
	index := uint16(row)*VGA_WIDTH + uint16(column)
	addr := (*uint16)(unsafe.Pointer(buffer + 2*uintptr(index)))
	*addr = makeVGAEntry(c, color)
}

func terminalPutChar(c byte) {
	terminalPutEntryAt(c, color, column, row)
	column++
	if column > VGA_WIDTH {
		column = 0
		row++
		if row > VGA_HEIGHT {
			row = 0
		}
	}
}

func writeString(data string) {
	row++
	column = 0
	for i := 0; i < len(data); i++ {
		terminalPutChar(data[i])
	}
}
