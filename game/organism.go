package game

type Position struct {
	X int
	Y int
}

type Organism struct {
	Size     int
	Energy   int
	Senses   []Sense
	Position Position
	Id       int
}
