using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("serialization");
$Go.import("ws_stream/serialization");

struct StructCapnproto {
	ts @0 :Int64;
	tag @1 :Text;
}
