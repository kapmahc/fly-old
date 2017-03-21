package auth

import (
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1/signatures"
	"github.com/SermoDigital/jose/jws"
	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/fly/web"
)

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"

	sendEmailJob = "auth.send-email"
)

// RegisterWorker register worker
func (p *Engine) RegisterWorker() {
	p.Server.RegisterTask(sendEmailJob, p.doSendEmail)
}

func (p *Engine) sendEmail(lng string, user *User, act string) {
	cm := jws.Claims{}
	cm.Set("act", act)
	cm.Set("uid", user.UID)
	tkn, err := p.Jwt.Sum(cm, time.Hour*6)
	if err != nil {
		log.Error(err)
		return
	}

	obj := struct {
		Home  string
		Token string
	}{
		Home:  web.Home(),
		Token: string(tkn),
	}
	subject, err := p.I18n.F(lng, fmt.Sprintf("auth.emails.%s.subject", act), obj)
	if err != nil {
		log.Error(err)
		return
	}
	body, err := p.I18n.F(lng, fmt.Sprintf("auth.emails.%s.body", act), obj)
	if err != nil {
		log.Error(err)
		return
	}

	// -----------------------
	task := signatures.TaskSignature{
		Name: sendEmailJob,
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "string",
				Value: user.Email,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: subject,
			},
			signatures.TaskArg{
				Type:  "string",
				Value: body,
			},
		},
	}

	if _, err := p.Server.SendTask(&task); err != nil {
		log.Error(err)
	}
}

func (p *Engine) parseToken(lng, tkn, act string) (*User, error) {
	cm, err := p.Jwt.Validate([]byte(tkn))
	if err != nil {
		return nil, err
	}
	if act != cm.Get("act").(string) {
		return nil, p.I18n.E(lng, "errors.bad-action")
	}
	return p.Dao.GetUserByUID(cm.Get("uid").(string))
}

func (p *Engine) doSendEmail(to, subject, body string) (interface{}, error) {
	if !web.IsProduction() {
		log.Infof("to %s\n%s\n%s", to, subject, body)
		return "done", nil
	}
	// TODO
	return "done", nil
}
