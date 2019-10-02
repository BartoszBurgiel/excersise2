package tourist

import (
	"math/rand"
	"time"
)

// Group represents a group of tourists
type Group struct {
	Tourists chan *Tourist

	//Every tourist used a computer
	IsDone chan bool

	// Count of the users that already were online
	UserCount int
}

// NewGroup constructor
func NewGroup() *Group {
	tourists := make(chan *Tourist, 25)

	// New seed
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 25; i++ {

		// Use random time as time user wants to spend online
		tourists <- &Tourist{i + 1, 0, rand.Intn(105) + 15}
	}

	return &Group{tourists, make(chan bool), 0}
}
