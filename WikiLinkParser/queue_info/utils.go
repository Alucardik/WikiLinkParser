package queue_info

import "github.com/streadway/amqp"

func declareQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

type ParseRequest struct {
	StartPage  string
	TargetPage string
	TaskId     uint64
}

type ParseResponse struct {
	TraceLen int
	Trace    string
	TaskId   uint64
}

type ParseMsg interface {
}
