package handle

type myItem interface {
	key() int
	name() string
	year() string
	rating() string
}

func (i Movie) key() int {
	return i.Key
}

func (i Movie) name() string {
	return i.Name
}

func (i Movie) year() string {
	return i.Year
}

func (i Movie) rating() string {
	return i.Rating
}

type myInterface interface {
	connectAndUpdateDB()
}
