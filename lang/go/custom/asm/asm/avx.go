package asm

import "unsafe"

//go:noescape
func _Avx2SsdInt16(a, b, c unsafe.Pointer)

//go:noescape
func _Avx2SsdInt16_2(a, b, c, d unsafe.Pointer)
