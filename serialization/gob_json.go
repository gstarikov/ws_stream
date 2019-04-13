package serialization

import (
	"encoding/gob"
	"io"
)

func GobJsonM(i interface{}, w io.Writer) error {
	dec := gob.NewEncoder(w)
	return dec.Encode(i)
}

func GobJsonU(r io.Reader, i interface{}) error {
	enc := gob.NewDecoder(r)
	return enc.Decode(i)
}
