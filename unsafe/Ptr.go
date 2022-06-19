package f64

import "unsafe"

// Casts uintptr into *float64
func NewPtr(uPtr uintptr) *float64 {
	return (*float64)(unsafe.Pointer(uPtr))
}
