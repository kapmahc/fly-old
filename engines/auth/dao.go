package auth

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/fly/engines/base"
)

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
	u.Password = base64.StdEncoding.EncodeToString(base.HmacSum([]byte(u.Password)))
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
func ConfirmUser(user uint) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(&User{}).Filter("id", user).Update(orm.Params{
		"confirmed_at": time.Now(),
		"updated_at":   time.Now(),
	})
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
		_, err = o.QueryTable(&p).Filter("id", p.ID).Update(orm.Params{
			"start_up":   begin,
			"shut_down":  end,
			"updated_at": time.Now(),
		})
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
