package serialization

import (
	"io"
)

func GogoprotoM(i interface{}, w io.Writer) error {
	str := i.(*TestSerializationStruct)
	globalGoGoProto.Ts = str.Ts
	globalGoGoProto.Tag = str.Tag

	//globalBuf.Reset()
	//globalBuf.Grow(globalGoGoProto.Size())
	//n, e := globalGoGoProto.MarshalTo(globalBuf.Bytes())
	b, e := globalGoGoProto.Marshal()
	if e != nil {
		return e
	}

	_, e = w.Write(b)
	//n2, e := globalBuf.WriteTo(w)
	//log.Printf("msg=[%s] n2=%d\n",b,n2)
	return e
}

func GogoprotoU(r io.Reader, i interface{}) error {
	str := i.(*TestSerializationStruct)
	globalBuf.Reset()
	globalBuf.ReadFrom(r)

	e := globalGoGoProto.Unmarshal(globalBuf.Bytes())
	if e != nil {
		return e
	}
	str.Tag = globalGoGoProto.Tag
	str.Ts = globalGoGoProto.Ts
	return nil
}

