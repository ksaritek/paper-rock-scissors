package model

type Session struct {
	Id       string `json:"id"`
	PlayerId string `json:"player_id"`
	Wins     int32  `json:"wins"`
	Losses   int32  `json:"losses"`
	Draws    int32  `json:"draws"`
}
