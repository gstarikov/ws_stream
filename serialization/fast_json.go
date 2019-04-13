package serialization

import (
	"github.com/intel-go/fastjson"
	"io"
)

func FastJsonM(i interface{}, w io.Writer) error {
	enc := fastjson.NewEncoder(w)
	return enc.Encode(i)
}

func FastJsonU(r io.Reader, i interface{}) error {
	dec := fastjson.NewDecoder(r)
	return dec.Decode(i)
}
