package fly

// Queue queue
type Queue interface {
	Send(name string, task []byte) error
	Receive(name string) ([]byte, error)
}

// Worker background worker
type Worker struct {
	Queue Queue `inject:""`
	Coder Coder `inject:""`
}

// Put put task into queue
func (p *Worker) Put(queue string, task interface{}) error {
	buf, err := p.Coder.Marshal(task)
	if err != nil {
		return err
	}
	return p.Queue.Send(queue, buf)
}

// Do process task queue
func (p *Worker) Do() {
	for {

	}
}
