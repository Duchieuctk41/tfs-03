package handle_goroutine

type Movier interface {
	key() int
	name() string
	year() string
	rating() string
}

func (m Movie) key() int {
	return m.Key
}

func (m Movie) name() string {
	return m.Name
}

func (m Movie) year() string {
	return m.Year
}

func (m Movie) rating() string {
	return m.Rating
}
