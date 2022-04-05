package worker

import (
	"regexp"
	"time"
)

// url formatting --------------------------------------------------------------

var RESTRICTED_PATTERNS = []*regexp.Regexp{
	// TODO: add js check
	// id reference withing the same page
	regexp.MustCompile(`^#.*`),
	// self-reference
	regexp.MustCompile(`^/?$`),
	// resource reference
	regexp.MustCompile(`^\./`),
	// adding blocked social media in light of the recent events
	regexp.MustCompile(`.*twitter.com.*`),
	regexp.MustCompile(`.*facebook.com.*`),
	regexp.MustCompile(`.*instagram.com.*`),
}

var FORMAT_PATTERNS = map[formatError]*regexp.Regexp{
	MISSING_PROTOCOL:         regexp.MustCompile(`^//`),
	MISSING_TOP_LEVEL_DOMAIN: regexp.MustCompile(`^/`),
	ID_REF:                   regexp.MustCompile(`#[^/]*$`),
	DOUBLE_QUOTES:            regexp.MustCompile(`^".*"`),
}

// parsing consts ----------------------------------------------------------------

const (
	MAX_DEPTH     = 5
	ROUTINE_LIM   = 1000
	REQUEST_DELAY = 5 * time.Second
	RETRY_CNT     = 2
)

// queue settings ----------------------------------------------------------------

const (
	SEND_QUEUE    = "worker_results"
	RECEIVE_QUEUE = "worker_tasks"
)
