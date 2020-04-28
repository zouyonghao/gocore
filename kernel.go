package kernel

func Main() {
	terminalInit()
	writeString("hello, kernel!")
	pmmInit()
	// for i := 0; i < 10; i++ {
	// 	writeString("a")
	// }
}
