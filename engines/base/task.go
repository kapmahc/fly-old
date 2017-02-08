package base

import (
	"fmt"
	"os"
	"time"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/signatures"
	"github.com/astaxie/beego"
)

var server *machinery.Server

func init() {
	cfg := config.Config{
		Broker:          beego.AppConfig.String("messagebroker"),
		ResultBackend:   beego.AppConfig.String("messagebroker"),
		DefaultQueue:    "default",
		ResultsExpireIn: int(time.Hour * 24 * 30 / time.Second),
	}
	var err error
	server, err = machinery.NewServer(&cfg)
	if err != nil {
		beego.Error(err)
	}

}

// SendTask send task
func SendTask(name string, args ...signatures.TaskArg) {
	server.SendTask(&signatures.TaskSignature{
		Name: name,
		Args: args,
	})
}

// RegisterTask resister task handler
func RegisterTask(name string, task interface{}) {
	server.RegisterTask(name, task)
}

// Worker worker start
func Worker() {
	name, err := os.Hostname()
	if err != nil {
		beego.Error(err)
	}
	beego.Info("start worker")
	worker := server.NewWorker(fmt.Sprintf("%s-worker", name))
	if err := worker.Launch(); err != nil {
		beego.Error(err)
	}
}
