package handle_goroutine

type Movie struct {
	Key    int
	Name   string `json:"name"`
	Year   string `json:"year"`
	Rating string `json:"rating"`
}

type ListMovie struct {
	Movies []Movie
}
