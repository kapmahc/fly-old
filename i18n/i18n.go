package i18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"golang.org/x/text/language"

	"github.com/kapmahc/fly"
)

// EXT locale file's ext
const EXT = ".json"

// Store locale store
type Store interface {
	Get(lang, code string) (string, error)
	Set(lang, code, message string) error
	Delete(lang, code string) error
	All(lang string) (map[string]string, error)
}

// I18N i18n
type I18N struct {
	Cache  fly.Cache  `inject:""`
	Store  Store      `inject:""`
	Logger fly.Logger `inject:""`
}

// Generate generate locale file
func (p *I18N) Generate(root string, lang language.Tag, val map[string]interface{}) error {
	os.MkdirAll(root, 0700)
	fd, err := os.OpenFile(
		path.Join(root, lang.String()+EXT),
		os.O_WRONLY|os.O_CREATE|os.O_EXCL,
		0600,
	)
	if err != nil {
		return err
	}
	defer fd.Close()
	enc := json.NewEncoder(fd)
	enc.SetIndent("", "  ")
	return enc.Encode(val)
}

// Load load from files
func (p *I18N) Load(root string) (map[string]map[string]string, error) {
	items := make(map[string]map[string]string)
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if filepath.Ext(name) == EXT {
			p.Logger.Info("find locale file", path)
		}

		fd, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fd.Close()

		var val interface{}
		dec := json.NewDecoder(fd)
		if err = dec.Decode(&val); err != nil {
			return err
		}

		tmp := make(map[string]string)
		err = p.parse("", tmp, val)
		if err != nil {
			return err
		}
		items[name[:len(name)-len(EXT)]] = tmp
		return nil
	}); err != nil {
		return nil, err
	}
	return items, nil
}

func (p *I18N) parse(prefix string, val map[string]string, data interface{}) error {

	switch data.(type) {
	case string:
		p.Logger.Debug(data)
		val[prefix] = data.(string)
	case map[string]interface{}:
		for k, v := range data.(map[string]interface{}) {
			if prefix != "" {
				k = prefix + "." + k
			}
			if err := p.parse(k, val, v); err != nil {
				return err
			}
		}
	default:
		p.Logger.Warn("ingnore %+v", data)
	}
	return nil
}

// T translate
func (p *I18N) T(lang, code string, args ...interface{}) string {
	msg, err := p.t(lang, code)
	if err != nil {
		return fmt.Sprintf("%s.%s", lang, code)
	}
	return fmt.Sprintf(msg, args...)
}

func (p *I18N) t(lang, code string) (msg string, err error) {
	key := fmt.Sprintf("locales/%s/%s", lang, code)
	if err = p.Cache.Get(key, &msg); err == nil {
		return
	}

	msg, err = p.Store.Get(lang, code)
	if err == nil {
		p.Cache.Set(key, msg, time.Hour*24)
		return
	}

	err = errors.New("not found")
	return
}
