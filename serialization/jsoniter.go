package serialization

import (
	"github.com/json-iterator/go"
	"io"
)

func JsoniterM(i interface{}, w io.Writer) error {
	dec := jsoniter.NewEncoder(w)
	return dec.Encode(i)
}

func JsoniterU(r io.Reader, i interface{}) error {
	enc := jsoniter.NewDecoder(r)
	return enc.Decode(i)
}
