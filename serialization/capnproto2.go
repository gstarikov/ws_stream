package serialization

import (
	"io"
	"zombiezen.com/go/capnproto2"
)

func CapnprotoM(i interface{}, w io.Writer) error {
	s := i.(*TestSerializationStruct)

	globalBuf.Reset()
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(globalBuf.Bytes()))
	if err != nil {
		return err
	}

	// Create a new struct.  Every message must have a root struct.
	sc, err := NewRootStructCapnproto(seg)
	if err != nil {
		return err
	}

	sc.SetTs(s.Ts)
	if err = sc.SetTag(s.Tag); err != nil {
		return err
	}

	return capnp.NewEncoder(w).Encode(msg)
}

func CapnprotoU(r io.Reader, i interface{}) error {

	//// Read the message from stdin.
	msg, err := capnp.NewDecoder(r).Decode()
	if err != nil {
		return err
	}

	// Extract the root struct from the message.
	rootStruct, err := ReadRootStructCapnproto(msg)
	if err != nil {
		return err
	}

	// Access fields from the struct.
	sOut := i.(*TestSerializationStruct)
	sOut.Ts = rootStruct.Ts()
	sOut.Tag, err = rootStruct.Tag()
	return err
}