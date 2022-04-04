package worker

import (
	"fmt"
	"sync"
)

// TODO: add methods and info to Launch method

type RequestStatus interface {
	IsFulfilled() bool
	Fulfill(foundTrace *trace)
	AddWorker()
	RemoveWorker()
	Await()
	Report()
}

type requestStatus struct {
	fulfilled bool
	path      trace
	lock      sync.RWMutex
	asscGroup sync.WaitGroup
}

func (rs *requestStatus) IsFulfilled() bool {
	rs.lock.RLock()
	curVal := rs.fulfilled
	rs.lock.RUnlock()

	return curVal
}

func (rs *requestStatus) Fulfill(foundTrace *trace) {
	rs.lock.Lock()
	rs.fulfilled = true
	foundTrace.copy(&rs.path)
	rs.lock.Unlock()
}

func (rs *requestStatus) AddWorker() {
	rs.asscGroup.Add(1)
}

func (rs *requestStatus) RemoveWorker() {
	rs.asscGroup.Done()
}

func (rs *requestStatus) Await() {
	rs.asscGroup.Wait()
}

func (rs *requestStatus) Report() {
	if rs.IsFulfilled() {
		fmt.Println("Request has been fulfilled")
		rs.path.formatAndPrint()
	} else {
		fmt.Println("Request has not been fulfilled")
	}
}
