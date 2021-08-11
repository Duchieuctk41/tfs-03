package handle

type Movie struct {
	Key int // add field
	Name string `json:"name"`
	Year string `json:"year"`
	Rating string `json:"rating"`
}

type ListMovie struct {
	Items []Movie
}
