package greeting

import (
    "bytes"
    "testing"
)

func TestSayHello(t *testing.T) {
    var buf bytes.Buffer
    SayHello(&buf)
    got := buf.String()
    want := "Hello, world!\n"
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
