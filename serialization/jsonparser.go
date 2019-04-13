package serialization

import (
	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"io"
)

func JsonparserM(i interface{}, w io.Writer) error {
	dec := jsoniter.NewEncoder(w)
	return dec.Encode(i)
}

func JsonparserU(r io.Reader, i interface{}) error {
	var paths = [][]string{
		[]string{"Ts"},
		[]string{"Tag"},
	}

	str := i.(*TestSerializationStruct)
	globalBuf.Reset()
	globalBuf.ReadFrom(r)
	jsonparser.EachKey(globalBuf.Bytes(), func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		switch idx {
		case 0:
			str.Ts, err = jsonparser.ParseInt(value)
		case 1:
			str.Tag, err = jsonparser.ParseString(value)
		}
	},
		paths...)

	return nil
}
