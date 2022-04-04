package limiter

type CountingSemaphore interface {
	SetTokenLim(int)
	GetToken()
	ReleaseToken()
}
