package worker

import (
	"WikiLinkParser/config"
	"WikiLinkParser/limiter"
	"WikiLinkParser/queue_info"
	"sync"
)

func Run() {
	forever := make(chan bool)

	qInfo := queue_info.CreateQueue()
	qInfo.InitQueueConnection(config.QUEUE_HOST, config.QUEUE_PORT, SEND_QUEUE, RECEIVE_QUEUE)

	defer qInfo.AbortConnection()

	msgHandler := func(msg queue_info.ParseMsg) {
		switch t := msg.(type) {
		case queue_info.ParseRequest:
			reqStat := launchWorker(t.StartPage, t.TargetPage)
			if reqStat.IsFulfilled() {
				qInfo.PublishTask(queue_info.ParseResponse{
					TraceLen: reqStat.GetTraceLen(),
					Trace:    reqStat.GetTrace(),
					TaskId:   t.TaskId,
				})
			} else {
				qInfo.PublishTask(queue_info.ParseResponse{
					TraceLen: 0,
					Trace:    "Request has failed",
					TaskId:   t.TaskId,
				})
			}
		}
	}

	go qInfo.GetResults(&msgHandler, queue_info.ParseRequest{})

	<-forever
}

func launchWorker(initUrl, targetUrl string) RequestStatus {
	var semaphore limiter.CountingSemaphore = &limiter.CountingSemaphoreImpl{}
	semaphore.SetTokenLim(ROUTINE_LIM)
	var newRequest RequestStatus = &requestStatus{
		fulfilled: false,
		path:      make(trace),
		lock:      sync.RWMutex{},
		asscGroup: sync.WaitGroup{},
	}

	curPage, err := createPageInfo(initUrl, targetUrl, nil, semaphore, newRequest)

	// launching from the function itself to ensure, that there will be at least one entry in the wait group
	if err == nil {
		curPage.parse()
	}

	newRequest.Await()
	newRequest.Report()

	return newRequest
}
