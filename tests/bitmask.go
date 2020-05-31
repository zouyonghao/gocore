package main

import "fmt"

type Bits uint8

const (
	F0 Bits = 1 << iota
	F1
	F2
	F3
	F4
	F5
	F6
	F7
)

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }

func main() {
	var b Bits
	b = Set(b, F0)
	b = Toggle(b, F2)
	for i, flag := range []Bits{F0, F1, F2} {
		fmt.Println(i, Has(b, flag))
	}
}
