package queue_info

import (
	"WikiLinkParser/error_utils"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type queueInfo struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	sendQueue  amqp.Queue
	rcvQueue   amqp.Queue
}

func (qi *queueInfo) InitQueueConnection(address string, port int, sendName, rcvName string) {
	conn, err := amqp.Dial("amqp://guest:guest@" + address + "/")
	qi.connection = conn
	error_utils.FailOnError("Error connecting to RabbitMQ", err)

	ch, err := conn.Channel()
	qi.channel = ch
	error_utils.FailOnError("Failed to open a channel", err)

	sendQ, err := declareQueue(ch, sendName)
	qi.sendQueue = sendQ
	error_utils.FailOnError("Failed to declare a queue", err)

	rcvQ, err := declareQueue(ch, rcvName)
	qi.rcvQueue = rcvQ
	error_utils.FailOnError("Failed to declare a queue", err)

	//qi.channel.NotifyReturn()
}

func (qi *queueInfo) AbortConnection() {
	if qi.channel != nil {
		qi.channel.Close()
	}

	if qi.connection != nil {
		qi.connection.Close()
	}
}

func (qi *queueInfo) PublishTask(msg ParseMsg) {
	workload, err := json.Marshal(msg)
	error_utils.NotifyOnError("Message cannot be encoded in json: %v\n", err)
	if err != nil {
		return
	}

	err = qi.channel.Publish(
		"",
		qi.sendQueue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  MSG_FORMAT,
			Body:         workload,
		})
	error_utils.NotifyOnError("Failed to publish a message", err)
}

func (qi *queueInfo) GetResults(handler *func(msg ParseMsg), msgType ParseMsg) {
	msgs, err := qi.channel.Consume(
		qi.rcvQueue.Name, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	error_utils.FailOnError("Failed to consume msgs from channel", err)

	var workloadResponse ParseResponse
	var workloadRequest ParseRequest
	var workload ParseMsg

	fmt.Println("Awaiting tasks...")
	for msg := range msgs {
		switch msgType.(type) {
		case ParseRequest:
			err = json.Unmarshal(msg.Body, &workloadRequest)
			workload = workloadRequest
		case ParseResponse:
			err = json.Unmarshal(msg.Body, &workloadResponse)
			workload = workloadResponse
		}
		error_utils.NotifyOnError("Can't decode json", err)

		if err == nil {
			fmt.Println("GOT TASK", workload)
			(*handler)(workload)
		}
	}
}
