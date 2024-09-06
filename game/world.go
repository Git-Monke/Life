package game

type World struct {
	tiles     [100][100]int
	organisms []Organism
}

func NewWorld() World {
	var grid [100][100]int
	var organisms []Organism

	return World{
		tiles:     grid,
		organisms: organisms,
	}
}
