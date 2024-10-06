package enums

type Strategy int

const (
	Algorithm Strategy = iota
	TOKEN_BUCKET
	SLIDING_WINDOW_LOG
)
