package i18n

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
)

// I18n i18n
type I18n struct {
	s Store
}

// New new  i18n
func New(s Store) *I18n {
	return &I18n{s: s}
}

// F format message
func (p *I18n) F(lng, code string, obj interface{}) (string, error) {
	msg, err := p.s.Get(lng, code)
	if err != nil {
		return "", err
	}
	tpl, err := template.New("").Parse(msg)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, obj)
	return buf.String(), err
}

//E create an i18n error
func (p *I18n) E(lang string, code string, args ...interface{}) error {
	msg, err := p.s.Get(lang, code)
	if err != nil {
		return errors.New(code)
	}
	return fmt.Errorf(msg, args...)
}

//T translate by lang tag
func (p *I18n) T(lng string, code string, args ...interface{}) string {
	msg, err := p.s.Get(lng, code)
	if err != nil {
		return code
	}
	return fmt.Sprintf(msg, args...)
}
