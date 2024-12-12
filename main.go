package main

import (
	"bytes"
	"fmt"

	"github.com/snowmerak/baseN/base"
)

func main() {
	bs, err := base.New("?!^#$")
	if err != nil {
		panic(err)
	}

	origin := []byte{0b10010010, 0b00100010, 0b10010011}

	encoder := bs.NewEncoder(bytes.NewReader(origin))
	res, err := encoder.Encode()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	buf := bytes.NewBuffer(nil)
	decoder := bs.NewDecoder(buf)
	decoded, err := decoder.Decode(res)
	if err != nil {
		panic(err)
	}

	fmt.Printf("origin: %08b\n", origin)
	fmt.Printf("decoded: %08b\n", decoded)
}
