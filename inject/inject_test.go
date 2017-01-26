package inject_test

import (
	"log"
	"testing"

	"github.com/kapmahc/fly/inject"
)

type S struct {
	S   string
	SSS string  `inject:"val.string"`
	I   int64   `inject:"val.int"`
	F   float32 `inject:"val.float"`
	SS  string
	F1  float32
	S1  *S1 `inject:""`
	S2  *S2 `inject:""`
	S3  *S1
	I4  I1
}

type S1 struct {
	S   string  `inject:"val.string"`
	I   int64   `inject:"val.int"`
	F   float32 `inject:"val.float"`
	S2  *S2     `inject:""`
	S22 *S2
}

type S2 struct {
	S   string
	SI1 I1 `inject:"is1"`
	SI2 I1 `inject:""`
}

type I1 interface {
	P()
}

type IS1 struct {
	S  string
	S2 string `inject:"val.string"`
}

func (p *IS1) P() {
	log.Println("is1")
}

type IS2 struct {
	S string `inject:"val.string"`
}

func (p *IS2) P() {
	log.Println("--- is2")
}
func (p *IS2) P1(s string) {
	log.Println("--- is2" + s)
}

func TestInject(t *testing.T) {
	inj := inject.New()
	inj.Debug(true)
	inj.MapTo("val.string", "hello, inject")
	inj.MapTo("val.float", float32(1.2))
	inj.MapTo("val.int", int64(124))
	inj.MapTo("is1", &IS1{})
	inj.Map(&S{}, &IS2{}, &IS1{})
	if err := inj.Populate(); err != nil {
		t.Fatal(err)
	}

	inj.Walk(func(o *inject.Object) error {
		t.Logf("%s: %+v", o.Name, o.Value)
		return nil
	})

	if _, err := inj.Invoke(func(i I1, s1 *IS1, s2 *IS2) {
		t.Log("-------------------------")
		i.P()
		s1.P()
		s2.P()
		t.Log(s2.S)
		s2.P1("aaa")
		t.Log("-------------------------")
	}); err != nil {
		t.Fatal(err)
	}
}
