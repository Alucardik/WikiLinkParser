package server

import (
	"WikiLinkParser/config"
	"WikiLinkParser/queue_info"
	"fmt"
)

func Run() {
	forever := make(chan bool)

	qInfo := queue_info.CreateQueue()
	qInfo.InitQueueConnection(config.QUEUE_HOST, config.QUEUE_PORT, SEND_QUEUE, RECEIVE_QUEUE)

	defer qInfo.AbortConnection()

	// TODO: add client handling via gRPC
	for i := 0; i < 5; i++ {
		qInfo.PublishTask(queue_info.ParseRequest{
			StartPage:  "https://en.wikipedia.org/wiki/Powder_House_Island",
			TargetPage: "https://en.wikipedia.org/wiki/Pointe_Mouillee_State_Game_Area",
		})
	}

	msgHandler := func(msg queue_info.ParseMsg) {
		switch t := msg.(type) {
		case queue_info.ParseResponse:
			fmt.Println(t.TraceLen)
			fmt.Println(t.Trace)
		}
	}

	go qInfo.GetResults(&msgHandler, queue_info.ParseResponse{})

	<-forever
}
