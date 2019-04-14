package serialization

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	log.SetFlags(log.Flags() | log.Lshortfile)
	os.Exit(m.Run())
}

func BenchmarkStdJson(b *testing.B) {
	BExamine(b, StdJsonM, StdJsonU)
}

func BenchmarkFastJson(b *testing.B) {
	BExamine(b, FastJsonM, FastJsonU)
}

func BenchmarkGobJson(b *testing.B) {
	BExamine(b, GobJsonM, GobJsonU)
}

func BenchmarkJsoniter(b *testing.B) {
	BExamine(b, JsoniterM, JsoniterU)
}

func BenchmarkSimpleCopy(b *testing.B) {
	BExamine(b, CopyM, CopyU)
}

func BenchmarkJsonparser(b *testing.B) {
	BExamine(b, JsonparserM, JsonparserU)
}

func BenchmarkEasyjson(b *testing.B) {
	BExamine(b, EasyjsonM, EasyjsonU)
}

func BenchmarkCapnproto(b *testing.B) {
	BExamine(b, CapnprotoM, CapnprotoU)
}

func BenchmarkGogoproto(b *testing.B) {
	BExamine(b, GogoprotoM, GogoprotoU)
}
