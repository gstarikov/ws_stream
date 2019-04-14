#!/usr/bin/env bash
#capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -ogo ./serialization/capnproto2.capnp
protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --gogofaster_out=. ./serialization/gogoproto.proto
