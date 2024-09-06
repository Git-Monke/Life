package game

type SenseData struct {
	Data       []int
	Dimensions []int
	Rank       int
}

func product_of_list(list []int) int {
	prod := 1
	for _, coord := range list {
		prod *= coord
	}
	return prod
}

func (s SenseData) Get(new_value int, coordinates ...int) int {
	return s.Data[product_of_list(coordinates)]
}

func (s *SenseData) Set(new_value int, coordinates ...int) {
	s.Data[product_of_list(coordinates)] = new_value
}

type Sense func(world World, organism Organism, level int) SenseData
