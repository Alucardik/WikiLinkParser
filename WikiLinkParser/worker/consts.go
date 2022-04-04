package worker

// html info -----------------------------------------------------------

const (
	ANCHOR       = "a"
	HREF         = "href"
	STD_PROTOCOL = "https:"
)

// format statuses ------------------------------------------------------

type formatError uint16

const (
	MISSING_PROTOCOL formatError = iota
	MISSING_TOP_LEVEL_DOMAIN
	ID_REF
	DOUBLE_QUOTES
)

// HTTP consts ------------------------------------------------------

const (
	TOO_MANY_REQUESTS = 429
)
