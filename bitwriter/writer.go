package bitwriter

const (
	mask1 = uint8(0b10000000)
	mask2 = uint8(0b11000000)
	mask3 = uint8(0b11100000)
	mask4 = uint8(0b11110000)
	mask5 = uint8(0b11111000)
	mask6 = uint8(0b11111100)
	mask7 = uint8(0b11111110)
	mask8 = uint8(0b11111111)
)

type Writer struct {
	buffer    []byte
	offset    int64
	bitOffset int64
}

func New() *Writer {
	return &Writer{buffer: []byte{0}, bitOffset: 8, offset: 0}
}

func (w *Writer) WriteBit(bit bool) {
	if w.bitOffset == 0 {
		w.buffer = append(w.buffer, 0)
		w.bitOffset = 8
		w.offset++
	}

	if bit {
		w.buffer[w.offset] |= 1 << uint(w.bitOffset-1)
	}

	w.bitOffset--
}

func (w *Writer) WriteByte(b byte) {
	if w.bitOffset == 0 {
		w.buffer = append(w.buffer, 0)
		w.bitOffset = 8
		w.offset++
	}

	w.buffer[w.offset] |= b >> uint(8-w.bitOffset)
	w.buffer = append(w.buffer, b<<uint(w.bitOffset))
	w.offset++
}

func (w *Writer) Bytes() []byte {
	cloned := make([]byte, len(w.buffer))
	copy(cloned, w.buffer)
	return cloned
}
