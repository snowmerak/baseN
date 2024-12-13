package base_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/snowmerak/baseN/base"
)

func BenchmarkEncoder_Encode(b *testing.B) {
	bs, err := base.New("?!^#$")
	if err != nil {
		b.Fatal(err)
	}

	buf := make([]byte, 1024)
	if _, err := rand.Read(buf); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = bs.NewEncoder(bytes.NewReader(buf)).Encode()
	}
}

func BenchmarkDecoder_Decode(b *testing.B) {
	bs, err := base.New("?!^#$")
	if err != nil {
		b.Fatal(err)
	}

	buf := make([]byte, 1024)
	if _, err := rand.Read(buf); err != nil {
		b.Fatal(err)
	}

	encoded, err := bs.NewEncoder(bytes.NewReader(buf)).Encode()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(nil)
		_, _ = bs.NewDecoder(buf).Decode(encoded)
	}
}
