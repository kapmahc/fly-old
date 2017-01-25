package fly_test

import (
	"testing"
	"time"

	"github.com/kapmahc/fly"
)

type M struct {
	I int
	F float32
	S string
	D time.Time
}

func TestSerial(t *testing.T) {
	m := M{I: 100, F: 1.2, S: "hello", D: time.Now()}
	b, e := fly.Marshal(&m)
	if e != nil {
		t.Fatal(e)
	}
	var tmp M
	if e = fly.Unmarshal(b, &tmp); e != nil {
		t.Fatal(e)
	}
	t.Logf("want %+v, get %+v", m, tmp)
	if m.I != tmp.I || m.S != tmp.S || m.F != tmp.F || m.D.UnixNano() != tmp.D.UnixNano() {
		t.Fatal("not equal")
	}
}
