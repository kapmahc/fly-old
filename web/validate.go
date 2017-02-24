package web

import (
	"net/http"

	"github.com/go-playground/form"
	validator "gopkg.in/go-playground/validator.v9"
)

// Validator valdator
type Validator struct {
	Validate *validator.Validate `inject:""`
	Decoder  *form.Decoder       `inject:""`
}

func (p *Validator) Bind(fm interface{}, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	if err := p.Decoder.Decode(fm, r.Form); err != nil {
		return err
	}
	return p.Validate.Struct(fm)
}
