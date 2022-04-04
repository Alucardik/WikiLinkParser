package limiter

type CountingSemaphoreImpl struct {
	Tokens chan struct{}
}

func (cs *CountingSemaphoreImpl) SetTokenLim(lim int) {
	cs.Tokens = make(chan struct{}, lim)
}

func (cs *CountingSemaphoreImpl) GetToken() {
	cs.Tokens <- struct{}{}
}

func (cs *CountingSemaphoreImpl) ReleaseToken() {
	<-cs.Tokens
}
