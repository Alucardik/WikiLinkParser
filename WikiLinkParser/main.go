package main

import (
	"WikiLinkParser/worker"
)

func main() {
	initPage := "https://en.wikipedia.org/wiki/Main_Page"
	targPage := "https://en.wikipedia.org/wiki/Norsemen_(TV_series)"
	worker.LaunchWorker(initPage, targPage)
}
