package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
	"github.com/kapmahc/fly/engines/base"
)

const (
	// RoleAdmin admin role
	RoleAdmin = "admin"
	// RoleRoot root role
	RoleRoot = "root"
	// UserTypeEmail email user
	UserTypeEmail = "email"

	// DefaultResourceType default resource type
	DefaultResourceType = "-"
	// DefaultResourceID default resourc id
	DefaultResourceID = 0
	// DefaultIP default ip
	DefaultIP = "0.0.0.0"
)

// User user
type User struct {
	base.Model

	Name            string     `json:"name"`
	Email           string     `json:"email"`
	UID             string     `json:"uid" orm:"column(uid)"`
	Password        string     `json:"-"`
	ProviderID      string     `json:"-" orm:"column(provider_id)"`
	ProviderType    string     `json:"providerNype"`
	Home            string     `json:"home"`
	Logo            string     `json:"logo"`
	SignInCount     uint       `json:"signInCount"`
	LastSignInAt    *time.Time `json:"lastSignInAt"`
	LastSignInIP    string     `json:"lastSignInIp" orm:"column(last_sign_in_ip)"`
	CurrentSignInAt *time.Time `json:"currentSignInAt"`
	CurrentSignInIP string     `json:"currentSignInIp" orm:"column(current_sign_in_ip)"`
	ConfirmedAt     *time.Time `json:"confirmedAt"`
	LockedAt        *time.Time `json:"lockedAt"`

	Logs []*Log `orm:"reverse(many)" json:"-"`
}

// TableName table name
func (User) TableName() string {
	return "users"
}

// IsConfirm is confirm?
func (p *User) IsConfirm() bool {
	return p.ConfirmedAt != nil
}

// IsLock is lock?
func (p *User) IsLock() bool {
	return p.LockedAt != nil
}

//SetGravatarLogo set logo by gravatar
func (p *User) SetGravatarLogo() {
	buf := md5.Sum([]byte(strings.ToLower(p.Email)))
	p.Logo = fmt.Sprintf("https://gravatar.com/avatar/%s.png", hex.EncodeToString(buf[:]))
}

//SetUID generate uid
func (p *User) SetUID() {
	p.UID = uuid.New().String()
}

func (p User) String() string {
	return fmt.Sprintf("%s<%s>", p.Name, p.Email)
}

// Attachment attachment
type Attachment struct {
	ID           uint
	Title        string
	URL          string
	Length       uint
	MediaType    string
	ResourceType string
	ResourceID   uint `orm:"column(resource_id)"`
	SortOrder    int
	CreatedAt    time.Time `orm:"auto_now_add"`

	User *User `orm:"rel(fk)"`
}

// TableName table name
func (Attachment) TableName() string {
	return "attachments"
}

// Log log
type Log struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt" orm:"auto_now_add"`
	IP        string    `json:"ip"`

	User *User `orm:"rel(fk)" json:"-"`
}

// TableName table name
func (Log) TableName() string {
	return "logs"
}

// Policy policy
type Policy struct {
	base.Model

	StartUp  time.Time
	ShutDown time.Time

	User *User `orm:"rel(fk)"`
	Role *Role `orm:"rel(fk)"`
}

//Enable is enable?
func (p *Policy) Enable() bool {
	now := time.Now()
	return now.After(p.StartUp) && now.Before(p.ShutDown)
}

// TableName table name
func (Policy) TableName() string {
	return "policies"
}

// Role role
type Role struct {
	base.Model

	Name         string
	ResourceID   uint `orm:"column(resource_id)"`
	ResourceType string
}

// TableName table name
func (Role) TableName() string {
	return "roles"
}

func (p Role) String() string {
	return fmt.Sprintf("%s@%s://%d", p.Name, p.ResourceType, p.ResourceID)
}

// Vote vote
type Vote struct {
	base.Model

	Point        int
	ResourceID   uint
	ResourceType string
}

// TableName table name
func (Vote) TableName() string {
	return "votes"
}

func init() {
	orm.RegisterModel(
		&User{}, &Log{}, &Role{}, &Policy{},
		&Attachment{},
		&Vote{},
	)
}
