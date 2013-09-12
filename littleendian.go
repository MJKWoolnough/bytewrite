package bytewrite

import "unsafe"

type littleEndian struct{}

func (e littleEndian) Float32(b []byte) float32 {
	i := e.Uint32(b)
	return *(*float32)(unsafe.Pointer(&i))
}

func (e littleEndian) PutFloat32Bits(f float32) []byte {
	return e.PutUint32(*(*uint32)(unsafe.Pointer(&f)))
}

func (e littleEndian) Float64(b []byte) float64 {
	i := e.Uint64(b)
	return *(*float64)(unsafe.Pointer(&i))
}

func (e littleEndian) PutFloat64(f float64) []byte {
	return e.PutUint64(*(*uint64)(unsafe.Pointer(&f)))
}

func (littleEndian) Uint16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func (littleEndian) PutUint16(v uint16) []byte {
	var b [2]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	return b[:]
}

func (littleEndian) Uint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func (littleEndian) PutUint32(v uint32) []byte {
	var b [4]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return b[:]
}

func (littleEndian) Uint64(b []byte) uint64 {
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func (littleEndian) PutUint64(v uint64) []byte {
	var b [8]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	return b[:]
}
