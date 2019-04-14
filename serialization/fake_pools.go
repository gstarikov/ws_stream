package serialization

import "bytes"

var (
	globalBuf = bytes.Buffer{}
	globalGoGoProto = StructProtobuf2{}
)
