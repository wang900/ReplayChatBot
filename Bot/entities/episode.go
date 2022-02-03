package entities

type Episode struct {
	ID    int  `json:"id"`
	Value int8 `json:"value"`
	Chat  Chat `json:"chat"`
}
