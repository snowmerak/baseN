package base

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/snowmerak/baseN/bitreader"
)

type Base struct {
	unit         int
	characterSet []byte
}

type CharacterSetLengthIsNotValidErr struct {
	Expected int
	Actual   int
}

func (e CharacterSetLengthIsNotValidErr) Error() string {
	return fmt.Sprintf("character set length is not valid, expected: %d, actual: %d", e.Expected, e.Actual)
}

func New(charSetParam string) (*Base, error) {
	charSet := []byte(charSetParam)
	unit := 1
	for i := 0; i < 64; i++ {
		if len(charSet) <= 1<<unit {
			break
		}
		unit++
	}
	return &Base{
		unit:         unit,
		characterSet: charSet,
	}, nil
}

func (b *Base) GetUnit() int {
	return b.unit
}

type Encoder struct {
	base   *Base
	reader io.Reader
}

func (b *Base) NewEncoder(reader io.Reader) *Encoder {
	return &Encoder{
		base:   b,
		reader: reader,
	}
}

func (e *Encoder) Encode() ([]byte, error) {
	res := make([]byte, 0)

	br, err := bitreader.New(e.reader)
	if err != nil {
		return nil, err
	}

	usePadding := false
	rm := 8 - (e.base.unit / 8)
	remains := e.base.unit % 8
	shifter := uint(8 - remains)
	if remains != 0 {
		rm--
	}
	if len(e.base.characterSet) < 1<<e.base.unit {
		usePadding = true
	}

	for {
		b, read, err := br.Read(int64(e.base.unit))
		if err != nil {
			break
		}

		if rm+len(b) < 8 {
			break
		}

		b = append(make([]byte, rm), b...)
		u := binary.BigEndian.Uint64(b)
		u >>= shifter + uint(e.base.unit) - uint(read)

		for u >= uint64(len(e.base.characterSet)) {
			res = append(res, e.base.characterSet[len(e.base.characterSet)-1])
			u -= uint64(len(e.base.characterSet)) - 1
		}

		res = append(res, e.base.characterSet[u])

		if u == uint64(len(e.base.characterSet))-1 && usePadding {
			res = append(res, e.base.characterSet[0])
		}
	}

	return res, nil
}
