package json

import (
	"encoding/json"
	"testing"
)

var (
	data = []byte(`{"a":1}`)
)

func BenchmarkJsonU(b *testing.B) {
	var v interface{}
	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &v)
	}
}

func BenchmarkJsonM(b *testing.B) {
	var v interface{}
	v = data
	for i := 0; i < b.N; i++ {
		json.Marshal(v)
	}
}
