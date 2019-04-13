package serialization

import (
	"io"
)

func EasyjsonM(i interface{}, w io.Writer) error {
	str := i.(*TestSerializationStruct)
	b, e := str.MarshalJSON()
	if e != nil {
		return e
	}
	_, e = w.Write(b)
	return e
}

func EasyjsonU(r io.Reader, i interface{}) error {
	str := i.(*TestSerializationStruct)
	globalBuf.Reset()
	globalBuf.ReadFrom(r)

	return str.UnmarshalJSON(globalBuf.Bytes())
}
