package main

// #cgo LDFLAGS: -L . -lstdc++
// #cgo CFLAGS: -I ./
// #include "foo.h"
import "C"

type GoFoo struct {
	foo C.Foo
}

func New() GoFoo {
	var ret GoFoo
	ret.foo = C.FooInit()
	return ret
}
func (f GoFoo) Free() {
	C.FooFree(f.foo)
}
func (f GoFoo) Bar() {
	C.FooBar(f.foo)
}

func main() {
	foo := New()
	foo.Bar()
	foo.Free()



	blockHeader := [32]C.uchar{
				0xd0, 0xda, 0xd7, 0x3f, 0xb2, 0xda, 0xbf, 0x33,
				0x53, 0xfd, 0xa1, 0x55, 0x71, 0xb4, 0xe5, 0xf6,
				0xac, 0x62, 0xff, 0x18, 0x7b, 0x35, 0x4f, 0xad,
				0xd4, 0x84, 0x0d, 0x9f, 0xf2, 0xf1, 0xaf, 0xdf,
			}

	seed := [32]C.uchar{
				0x07, 0x37, 0x52, 0x07, 0x81, 0x34, 0x5b, 0x11,
				0xb7, 0xbd, 0x0f, 0x84, 0x3c, 0x1b, 0xdd, 0x9a,
				0xea, 0x81, 0xb6, 0xda, 0x94, 0xfd, 0x14, 0x1c,
				0xc9, 0xf2, 0xdf, 0x53, 0xac, 0x67, 0x44, 0xd2,
			}

	hash := [32]C.uchar{
				0xe3, 0x5d, 0xa5, 0x47, 0x95, 0xd8, 0x2f, 0x85,
				0x49, 0xc0, 0xe5, 0x80, 0xcb, 0xf2, 0xe3, 0x75,
				0x7a, 0xb5, 0xef, 0x8f, 0xed, 0x1b, 0xdb, 0xe4,
				0x39, 0x41, 0x6c, 0x7e, 0x6f, 0x8d, 0xf2, 0x27,
			}

	C.get(&blockHeader[0])
	C.get(&seed[0])
	C.get(&hash[0])
}
