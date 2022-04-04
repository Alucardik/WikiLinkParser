package worker

import (
	"WikiLinkParser/limiter"
	"errors"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
)

type pageInfo struct {
	target    string
	current   string
	topDomain string
	urls      stringSet
	tr        trace
	semaphore limiter.CountingSemaphore
	status    RequestStatus
}

func (p *pageInfo) getUrls(resp *http.Response) {
	results := make(stringSet)
	tokenizer := html.NewTokenizer(resp.Body)

	// TODO: create a workaround initial value
	for tType := html.CommentToken; tType != html.ErrorToken; tType = tokenizer.Next() {
		if tType != html.StartTagToken {
			continue
		}

		t := tokenizer.Token()
		if t.Data != ANCHOR {
			continue
		}

		for _, attr := range t.Attr {
			if attr.Key == HREF {
				anchorVal := attr.Val
				if checkForRestricted(anchorVal) {
					anchorVal = formatUrl(anchorVal, p.topDomain)
					if !p.tr.contains(anchorVal) {
						results.add(anchorVal)
					}
				}
				break
			}
		}
	}

	p.urls = results
}

func (p *pageInfo) crawlUrls() {
	for nextTarget := range p.urls {
		if len(p.tr) < MAX_DEPTH {
			if p.status.IsFulfilled() {
				break
			}

			nextPage, err := createPageInfo(nextTarget, p.target, &p.tr, p.semaphore, p.status)
			if err == nil {
				go nextPage.parse()
			}
		}
	}
}

func (p *pageInfo) wrap() {
	p.semaphore.ReleaseToken()
	p.status.RemoveWorker()
}

func (p *pageInfo) parse() {
	p.status.AddWorker()
	p.semaphore.GetToken()

	if p.status.IsFulfilled() {
		p.wrap()
		return
	}

	if p.current == p.target {
		p.status.Fulfill(&p.tr)
		p.wrap()
		return
	}

	resp, ok := sendRequest(p.current)
	if !ok {
		p.wrap()
		return
	}

	p.getUrls(resp)
	p.crawlUrls()
	p.wrap()
}

func createPageInfo(
	curUrl, targetUrl string,
	prevTrace *trace,
	semaphore limiter.CountingSemaphore,
	status RequestStatus) (pageInfo, error) {
	urlInfo, err := url.Parse(curUrl)
	newTrace := make(trace)
	prevTrace.copy(&newTrace)
	newTrace.add(curUrl)

	if urlInfo == nil || err != nil {
		return pageInfo{}, errors.New("error parsing url")
	}

	return pageInfo{
		current:   curUrl,
		target:    targetUrl,
		topDomain: urlInfo.Host,
		tr:        newTrace,
		semaphore: semaphore,
		status:    status,
	}, nil
}
