package main

import (
	"github.com/posener/wstest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestWsHandler(t *testing.T) {
	var err error

	d := wstest.NewDialer(http.HandlerFunc(WsHandler))
	c, resp, err := d.Dial("ws://"+"localhost"+"/ws", nil)
	defer func() { assert.NoError(t, c.Close()) }()

	assert.NoError(t, err)

	if got, want := resp.StatusCode, http.StatusSwitchingProtocols; got != want {
		t.Fatalf("resp.StatusCode = %q, want %q", got, want)
	}

	reqStr := "test"
	assert.NoError(t, c.WriteJSON(reqStr))

	var respStr string
	assert.NoError(t, c.ReadJSON(&respStr))

	assert.EqualValues(t, reqStr, respStr)

	t.Log(respStr)

	assert.NoError(t, c.Close())

	time.Sleep(time.Second)
}
