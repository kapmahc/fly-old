package pos

import (
	"time"

	"github.com/kapmahc/fly/web"
)

// Company company
type Company struct {
	web.Model
	Name    string
	Tel     string
	Fax     string
	Owner   string
	Account string
	Address string
	Email   string
	WebSite string
}

// Department department
type Department struct {
	web.Model
	CompanyID uint
	Company   Company
	ManagerID uint
	Manager   []Employee
}

// Employee employee
type Employee struct {
	web.Model
	Name         string
	Sex          string
	Tel          string
	Email        string
	Address      string
	Birthday     time.Time
	Contacts     string
	DepartmentID uint
	Department   Department
}
