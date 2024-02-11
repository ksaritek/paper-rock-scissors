package model

type Choice int32

const (
	INVALID_CHOICE Choice = 0
	ROCK           Choice = 1
	PAPER          Choice = 2
	SCISSORS       Choice = 3
)

type Move struct {
	Id     string `json:"id"`
	Choice Choice `json:"choice"`
}
