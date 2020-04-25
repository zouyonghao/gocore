package runtime

import "unsafe"

// For gccgo, use go:linkname to rename compiler-called functions to
// themselves, so that the compiler will export them.
//
//go:linkname memequal runtime.memequal
//go:linkname memequal16 runtime.memequal16$descriptor
//go:linkname memequal32 runtime.memequal32$descriptor
//go:linkname memhash runtime.memhash
//go:linkname memhash16 runtime.memhash16$descriptor
//go:linkname memhash32 runtime.memhash32$descriptor

const (
	// Constants for multiplication: four random odd 32-bit numbers.
	m1 = 3168982561
	m2 = 3339683297
	m3 = 832293441
	m4 = 2336365089
)

var hashkey [4]uintptr

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func readUnaligned32(p unsafe.Pointer) uint32 {
	return *(*uint32)(p)
}

func readUnaligned64(p unsafe.Pointer) uint64 {
	return *(*uint64)(p)
}

func memhash(p unsafe.Pointer, seed, s uintptr) uintptr {

	h := uint32(seed + s*hashkey[0])
tail:
	switch {
	case s == 0:
	case s < 4:
		h ^= uint32(*(*byte)(p))
		h ^= uint32(*(*byte)(add(p, s>>1))) << 8
		h ^= uint32(*(*byte)(add(p, s-1))) << 16
		h = rotl_15(h*m1) * m2
	case s == 4:
		h ^= readUnaligned32(p)
		h = rotl_15(h*m1) * m2
	case s <= 8:
		h ^= readUnaligned32(p)
		h = rotl_15(h*m1) * m2
		h ^= readUnaligned32(add(p, s-4))
		h = rotl_15(h*m1) * m2
	case s <= 16:
		h ^= readUnaligned32(p)
		h = rotl_15(h*m1) * m2
		h ^= readUnaligned32(add(p, 4))
		h = rotl_15(h*m1) * m2
		h ^= readUnaligned32(add(p, s-8))
		h = rotl_15(h*m1) * m2
		h ^= readUnaligned32(add(p, s-4))
		h = rotl_15(h*m1) * m2
	default:
		v1 := h
		v2 := uint32(seed * hashkey[1])
		v3 := uint32(seed * hashkey[2])
		v4 := uint32(seed * hashkey[3])
		for s >= 16 {
			v1 ^= readUnaligned32(p)
			v1 = rotl_15(v1*m1) * m2
			p = add(p, 4)
			v2 ^= readUnaligned32(p)
			v2 = rotl_15(v2*m2) * m3
			p = add(p, 4)
			v3 ^= readUnaligned32(p)
			v3 = rotl_15(v3*m3) * m4
			p = add(p, 4)
			v4 ^= readUnaligned32(p)
			v4 = rotl_15(v4*m4) * m1
			p = add(p, 4)
			s -= 16
		}
		h = v1 ^ v2 ^ v3 ^ v4
		goto tail
	}
	h ^= h >> 17
	h *= m3
	h ^= h >> 13
	h *= m4
	h ^= h >> 16
	return uintptr(h)
}

func memhash16(p unsafe.Pointer, h uintptr) uintptr {
	return memhash(p, h, 2)
}

func memhash32(p unsafe.Pointer, seed uintptr) uintptr {
	h := uint32(seed + 4*hashkey[0])
	h ^= readUnaligned32(p)
	h = rotl_15(h*m1) * m2
	h ^= h >> 17
	h *= m3
	h ^= h >> 13
	h *= m4
	h ^= h >> 16
	return uintptr(h)
}

// Note: in order to get the compiler to issue rotl instructions, we
// need to constant fold the shift amount by hand.
// TODO: convince the compiler to issue rotl instructions after inlining.
func rotl_15(x uint32) uint32 {
	return (x << 15) | (x >> (32 - 15))
}

//extern __builtin_memcmp
func __builtin_memcmp(a, b unsafe.Pointer, size uintptr) bool

func memequal(a, b unsafe.Pointer, size uintptr) bool {
	return __builtin_memcmp(a, b, size)
}

func memequal16(p, q unsafe.Pointer) bool {
	return *(*int16)(p) == *(*int16)(q)
}

func memequal32(p, q unsafe.Pointer) bool {
	return *(*int32)(p) == *(*int32)(q)
}
