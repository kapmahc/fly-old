package migrate

// Dialect dialect
type Dialect interface {
	Up(file string) error
	Down(file string) error
	Exec(sql string) error
	Version() (*Model, error)
	All() ([]Model, error)
	Check()
}
