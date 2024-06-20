package bytewrite // import "vimagination.zapto.org/bytewrite"

var (
	// BigEndian implementinterface Endian.
	BigEndian bigEndian
	// LittleEndian implementinterface Endian.
	LittleEndian littleEndian
	_            Endian = BigEndian
	_            Endian = LittleEndian
)

// Endian represents the numerous methds available for easy endian conversion.
type Endian interface {
	// Takes a slice of bytes and returns a float32 oriented according to the endianness.
	Float32(b []byte) float32
	// Takes a float32 and returns a slice of bytes, ordered according to the endianness.
	PutFloat32(f float32) []byte
	// Takes a slice of bytes and returns a float64 oriented according to the endianness.
	Float64(b []byte) float64
	// Takes a float64 and returns a slice of bytes, ordered according to the endianness.
	PutFloat64(f float64) []byte
	// Takes a slice of bytes and returns a uint16 oriented according to the endianness.
	Uint16(b []byte) uint16
	// Takes a uint16 and returns a slice of bytes, ordered according to the endianness.
	PutUint16(v uint16) []byte
	// Takes a slice of bytes and returns a uint32 oriented according to the endianness.
	Uint32(b []byte) uint32
	// Takes a uint32 and returns a slice of bytes, ordered according to the endianness.
	PutUint32(v uint32) []byte
	// Takes a slice of bytes and returns a uint64 oriented according to the endianness.
	Uint64(b []byte) uint64
	// Takes a uint64 and returns a slice of bytes, ordered according to the endianness.
	PutUint64(v uint64) []byte
}
