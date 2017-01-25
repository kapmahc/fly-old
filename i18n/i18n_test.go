package i18n_test

import (
	"os"
	"testing"

	"github.com/kapmahc/fly"
	"github.com/kapmahc/fly/i18n"
	"golang.org/x/text/language"
)

func TestI18nLoad(t *testing.T) {
	const root = "locales"
	os.RemoveAll(root)
	in := i18n.I18N{Logger: &fly.ConsoleLogger{}}

	for _, lang := range []language.Tag{language.AmericanEnglish, language.SimplifiedChinese} {
		if err := in.Generate(root, lang, map[string]interface{}{
			"hi": "Hello!",
			"buttons": map[string]string{
				"submit": "submit",
				"reset":  "reset",
			},
			"forum": map[string]interface{}{
				"pages": map[string]string{
					"about": "About us",
					"home":  "Home",
				},
			},
		}); err != nil {
			t.Fatal(err)
		}
	}

	val, err := in.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", val)
}
