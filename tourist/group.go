package tourist

import (
	"math/rand"
	"time"
)

// Group represents a group of tourists
type Group struct {
	Tourists chan *Tourist

	//Every tourist used a computer
	IsDone bool

	TouristsToGo, TouristsDone int
}

// NewGroup constructor
func NewGroup() *Group {
	tourists := make(chan *Tourist, 25)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 25; i++ {
		tourists <- &Tourist{i + 1, 0, rand.Intn(105) + 15}
	}

	return &Group{tourists, false, 25, 0}
}
