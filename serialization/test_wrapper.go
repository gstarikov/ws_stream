package serialization

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

type Marshall func(interface{}, io.Writer) error
type Unmashall func(io.Reader, interface{}) error

func ExamineBench(t testing.TB, marshall Marshall, unmashall Unmashall, buf *bytes.Buffer, from, to *TestSerializationStruct) {
	if marshall(from, buf) != nil || unmashall(buf, to) != nil {
		t.FailNow()
	}
}

func BExamine(b *testing.B, marshall Marshall, unmashall Unmashall) {
	buf := &bytes.Buffer{}
	from := &TestSerializationStruct{
		Ts:  time.Now().UnixNano(),
		Tag: "Test Tag",
	}

	to := &TestSerializationStruct{} // to avoid allocation

	assert.Nil(b, marshall(from, buf))
	assert.Nil(b, unmashall(buf, to))
	assert.Equal(b, from, to)

	b.ReportAllocs()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		ExamineBench(b, marshall, unmashall, buf, from, to)
	}
}
