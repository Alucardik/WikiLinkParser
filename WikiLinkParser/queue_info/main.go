package queue_info

type QueueInfo interface {
	InitQueueConnection(address string, port int, sendName, rcvName string)
	AbortConnection()
	PublishTask(msg ParseMsg)
	GetResults(handler *func(msg ParseMsg), msgType ParseMsg)
}

func CreateQueue() QueueInfo {
	newQueueInfo := queueInfo{}
	return &newQueueInfo
}
