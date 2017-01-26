package migrate

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"time"
)

const (
	// FORMAT time format
	FORMAT = "20060102150405"
	// UP up
	UP = "--- UP ---"
	// END end
	END = "--- END ---"
	// DOWN down
	DOWN = "--- DOWN ---"
	// EXT file ext
	EXT = ".sql"
)

// Model migration
type Model struct {
	File      string    `gorm:"primary_key;size:255;unique"`
	Applied   bool      `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

// Up up scripts
func (p Model) Up() (string, error) {
	return p.read(UP)
}
func (p Model) read(start string) (string, error) {
	fd, err := os.Open(p.File)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	var buf bytes.Buffer
	begin := false
	san := bufio.NewScanner(fd)
	for san.Scan() {
		line := san.Text()

		if line == start {
			begin = true
			continue
		}
		if begin && line == END {
			break
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		if begin {
			buf.WriteString(line)
		}
	}

	return buf.String(), nil
}

// Down down scripts
func (p Model) Down() (string, error) {
	return p.read(DOWN)
}

// TableName table name
func (p Model) TableName() string {
	return "schema_migrations"
}
