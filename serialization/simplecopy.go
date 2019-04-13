package serialization

import (
	"io"
)

var ci *TestSerializationStruct

func CopyM(i interface{}, w io.Writer) error {
	ci = i.(*TestSerializationStruct)
	return nil
}

func CopyU(r io.Reader, i interface{}) error {
	*(i.(*TestSerializationStruct)) = *ci
	return nil
}
