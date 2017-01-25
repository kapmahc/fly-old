package fly

import (
	"fmt"
	"os"
)

// Queue queue
type Queue interface {
	Send(name string, task []byte) error
	Receive(names ...string) (string, []byte, error)
}

// Worker background worker
type Worker struct {
	Queue    Queue  `inject:""`
	Coder    Coder  `inject:""`
	Logger   Logger `inject:""`
	Handlers map[string]func([]byte) error
}

// Register register handler
func (p *Worker) Register(que string, hnd func([]byte) error) {
	if _, ok := p.Handlers[que]; ok {
		p.Logger.Warn("override handler for queue", que)
	}
	p.Handlers[que] = hnd
}

// Put put task into queue
func (p *Worker) Put(queue string, task interface{}) error {
	buf, err := p.Coder.Marshal(task)
	if err != nil {
		return err
	}
	return p.Queue.Send(queue, buf)
}

func (p *Worker) handle(names ...string) error {
	que, buf, err := p.Queue.Receive(names...)
	if err != nil {
		return err
	}
	hnd, ok := p.Handlers[que]
	if !ok {
		return fmt.Errorf("can't find handler for queue %s", que)
	}
	return hnd(buf)
}

// Do process task queue
func (p *Worker) Do() {
	p.Logger.Info("start worker...")
	var names []string
	for k := range p.Handlers {
		names = append(names, k)
	}
	if len(names) == 0 {
		p.Logger.Error("task handlers is empty")
		return
	}
	p.Logger.Info("listen to queues", names)

	for {
		const file = ".stop"
		if _, err := os.Stat(file); err == nil {
			p.Logger.Info("find file", file, ", exit")
			return
		}
		if err := p.handle(names...); err != nil {
			p.Logger.Error(err)
		}
	}
}
