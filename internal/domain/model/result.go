package model

type RoundResult int32

const (
	INVALID_RESULT RoundResult = 0
	WIN            RoundResult = 1
	LOSS           RoundResult = 2
	DRAW           RoundResult = 3
)
