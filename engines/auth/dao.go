package auth

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/fly/engines/base"
)

// SetSignIn set sign-in info
func SetSignIn(user *User, ip string) error {
	// ip, _, err := net.SplitHostPort(req.RemoteAddr)
	// if err != nil {
	// 	return err
	// }
	o := orm.NewOrm()
	user.LastSignInAt = user.CurrentSignInAt
	user.LastSignInIP = user.CurrentSignInIP
	user.CurrentSignInIP = ip
	n := time.Now()
	user.CurrentSignInAt = &n
	user.SignInCount++

	_, err := o.Update(
		user,
		"last_sign_in_at", "last_sign_in_ip",
		"current_sign_in_at", "current_sign_in_ip",
		"sign_in_count", "updated_at",
	)
	return err

}

// AddEmailUser add user
func AddEmailUser(name, email, password string) (*User, error) {
	var u User
	o := orm.NewOrm()
	count, err := o.QueryTable(&u).Filter("email", email).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("email %s already exists", email)
	}
	u.Email = email
	u.Name = name
	u.Password = base64.StdEncoding.EncodeToString(base.HmacSum([]byte(password)))
	u.ProviderID = email
	u.ProviderType = UserTypeEmail
	u.CurrentSignInIP = DefaultIP
	u.LastSignInIP = DefaultIP
	u.SetUID()
	u.SetGravatarLogo()
	_, err = o.Insert(&u)
	return &u, err
}

// ConfirmUser confirm
func ConfirmUser(user *User) error {
	n := time.Now()
	user.ConfirmedAt = &n
	o := orm.NewOrm()
	_, err := o.Update(user, "confirmed_at", "updated_at")
	return err
}

// Allow apply role to user
func Allow(user *User, role string, years, months, days int) error {
	r, err := getRole(role, DefaultResourceType, DefaultResourceID)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	var p Policy
	err = o.QueryTable(&p).
		Filter("user_id", user.ID).
		Filter("role_id", r.ID).One(&p)

	begin := time.Now()
	end := begin.AddDate(years, months, days)
	if err == nil {
		p.StartUp = begin
		p.ShutDown = end
		_, err = o.Update(&p, "start_up", "shut_down", "updated_at")
	} else if err == orm.ErrNoRows {
		p.StartUp = begin
		p.ShutDown = end
		p.Role = r
		p.User = user
		_, err = o.Insert(&p)
	}
	return err
}

// Deny deny role from user
func Deny(user *User, role string) error {
	r, err := getRole(role, DefaultResourceType, DefaultResourceID)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.QueryTable(&Policy{}).
		Filter("user_id", user.ID).
		Filter("role_id", r.ID).Delete()
	return err
}

func getRole(name, rty string, rid uint) (*Role, error) {
	o := orm.NewOrm()
	var r Role
	err := o.QueryTable(&r).
		Filter("name", name).
		Filter("resource_type", rty).
		Filter("resource_id", rid).
		One(&r)

	if err == orm.ErrNoRows {
		r.Name = name
		r.ResourceID = rid
		r.ResourceType = rty
		_, err = o.Insert(&r)
	}
	if err == nil {
		return &r, nil
	}
	return nil, err
}
