package web

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// NewQueue new queue
func NewQueue() *Queue {
	return &Queue{tasks: make(map[string]func([]byte) error)}
}

// Queue message queue
type Queue struct {
	tasks map[string]func([]byte) error
}

// Register register job
func (p *Queue) Register(n string, f func([]byte) error) {
	if _, ok := p.tasks[n]; ok {
		log.Warningf("task %s already exists!", n)
	}
	p.tasks[n] = f
}

// Do do job
func (p *Queue) Do(name string) error {
	log.Infof("waiting for messages, to exit press CTRL+C")
	return p.open(func(ch *amqp.Channel) error {
		if err := ch.Qos(1, 0, false); err != nil {
			return err
		}
		for n, f := range p.tasks {
			qu, err := ch.QueueDeclare(n, true, false, false, false, nil)
			if err != nil {
				return err
			}
			msgs, err := ch.Consume(qu.Name, name, false, false, false, false, nil)
			if err != nil {
				return err
			}
			for d := range msgs {
				log.Infof("receive task %s@%s", d.MessageId, qu.Name)
				d.Ack(false)
				if err := f(d.Body); err != nil {
					return err
				}
				log.Info("done")
			}
		}
		return nil
	})
}

// Send send job
func (p *Queue) Send(exchange, queue string, body interface{}) {
	if err := p.open(func(ch *amqp.Channel) error {
		qu, err := ch.QueueDeclare(queue, true, false, false, false, nil)
		if err != nil {
			return err
		}
		buf, err := json.Marshal(body)
		if err != nil {
			return err
		}
		return ch.Publish(exchange, qu.Name, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			MessageId:    uuid.New().String(),
			Body:         buf,
		})
	}); err != nil {
		log.Error(err)
	}
}

func (p *Queue) open(f func(*amqp.Channel) error) error {
	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		viper.GetString("rabbitmq.user"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.host"),
		viper.GetInt("rabbitmq.port"),
		viper.GetString("rabbitmq.virtual"),
	))
	if err != nil {
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	return f(ch)
}
