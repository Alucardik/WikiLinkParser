package worker

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// util functions -----------------------------------------------

func sendRequestRaw(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Encountered error: %v\n", err)
		return nil, err
	}

	return resp, nil
}

func sendRequest(url string) (*http.Response, bool) {
	for i := 1; i < RETRY_CNT; i++ {
		resp, err := sendRequestRaw(url)
		if err != nil {
			return nil, false
		}

		if resp.StatusCode == TOO_MANY_REQUESTS {
			time.Sleep(REQUEST_DELAY)
			continue
		} else if !checkResponse(resp.StatusCode) {
			log.Println("Non-ok response status: ", resp.Status)
			return nil, false
		}

		return resp, true
	}

	return nil, false
}

func checkResponse(respCode int) bool {
	return respCode/100 == 2
}

func checkForRestricted(testPattern string) bool {
	for _, restricted := range RESTRICTED_PATTERNS {
		if restricted.MatchString(testPattern) {
			//fmt.Println("Matched ", testPattern, " with ", restricted.String())
			return false
		}
	}

	return true
}

func resolveInconsistency(url, topDomain string, incType formatError) string {
	switch incType {
	case DOUBLE_QUOTES:
		url = strings.Trim(url, `"`)
	case ID_REF:
		i := len(url) - 1
		for ; url[i] != '#'; i-- {
		}

		return url[:i]
	case MISSING_PROTOCOL:
		return STD_PROTOCOL + url
	case MISSING_TOP_LEVEL_DOMAIN:
		return STD_PROTOCOL + "//" + topDomain + url
	}

	return url
}

func formatUrl(url, topDomain string) string {
	for errType, pattern := range FORMAT_PATTERNS {
		if pattern.MatchString(url) {
			url = resolveInconsistency(url, topDomain, errType)
		}
	}

	return url
}

// util types -------------------------------------------------------------------------

// trace represents already visited links on the search path, where values form chronological order
type trace map[string]uint64

func (t *trace) contains(val string) bool {
	_, ok := (*t)[val]
	return ok
}

func (t *trace) add(link string) {
	(*t)[link] = uint64(len(*t) + 1)
}

// copy contents of the caller trace to another trace
func (t *trace) copy(target *trace) {
	if t == nil {
		return
	}

	for k, v := range *t {
		(*target)[k] = v
	}
}

func (t *trace) formatAndPrint() {
	fmt.Printf("Trace length: %d\n-------------------------\n", len(*t))

	chronologicalOrd := make([]string, len(*t))
	for k, v := range *t {
		chronologicalOrd[v-1] = k
	}

	for i, v := range chronologicalOrd {
		fmt.Print(v)
		if i != len(chronologicalOrd)-1 {
			fmt.Print(" -->\n")
		}
	}
	fmt.Println()
}

// stringSet represents a set of strings
type stringSet map[string]bool

func (ss *stringSet) add(val string) {
	_, ok := (*ss)[val]
	if !ok {
		(*ss)[val] = true
	}
}
