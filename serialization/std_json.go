package serialization

import (
	"encoding/json"
	"io"
)

func StdJsonM(i interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(i)
}

func StdJsonU(r io.Reader, i interface{}) error {
	dec := json.NewDecoder(r)
	return dec.Decode(i)
}
