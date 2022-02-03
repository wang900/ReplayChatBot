package entities

type Anime struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Episodes []Episode `json:"episodes"`
}
