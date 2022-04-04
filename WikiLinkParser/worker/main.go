package worker

import (
	"WikiLinkParser/limiter"
	"sync"
)

func LaunchWorker(initUrl, targetUrl string) {
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
}
