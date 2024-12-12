package main

import (
	"bytes"
	"fmt"

	"github.com/snowmerak/baseN/base"
)

func main() {
	bs, err := base.New('?', '!')
	if err != nil {
		panic(err)
	}

	encoder := bs.NewEncoder(bytes.NewReader([]byte{0b10010010, 0b00100010, 0b10010011}))
	res, err := encoder.Encode()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
