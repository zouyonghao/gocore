package stdlib

import "video"

var ErrorMsg [13][40]byte

func CopyStr(array *[40]byte, str string) {
	if len(str) > 40 {
		video.Error(ErrorMsg[0], len(str), true)
	}
	for i := 0; i < len(str); i++ {
		array[i] = str[i]
	}
}

func initErrs() {
	CopyStr(&ErrorMsg[0], "Error message too long")
	CopyStr(&ErrorMsg[1], "Slice Index out of Bounds Exception")
	CopyStr(&ErrorMsg[2], "Array Index out of Bounds Exception")
	CopyStr(&ErrorMsg[3], "String Index out of Bounds Exception")
	CopyStr(&ErrorMsg[4], "Slice Slice out of Bounds Exception")
	CopyStr(&ErrorMsg[5], "Array Slice out of Bounds Exception")
	CopyStr(&ErrorMsg[6], "String Slice out of Bounds Exception")
	CopyStr(&ErrorMsg[7], "Nil Pointer Exception")
	CopyStr(&ErrorMsg[8], "Make slice out of bounds")
	CopyStr(&ErrorMsg[9], "Make map out of bounds")
	CopyStr(&ErrorMsg[10], "Make chan out of bounds")
	CopyStr(&ErrorMsg[11], "Division By Zero Exception")
	CopyStr(&ErrorMsg[12], "Unknown Exception")

}

func ErrCode(code int32) {
	if code > 10 {
		video.Error(ErrorMsg[12], int(code), true)
	}
	video.Error(ErrorMsg[code+1], int(code), true)
}
