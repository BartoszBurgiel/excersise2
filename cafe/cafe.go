package cafe

import (
	"fmt"
	"internetCafe/tourist"
)

// Cafe represents an internetcafe with its computers
type Cafe struct {
	Computers    [8]Computer
	FreeComputer chan bool
}

// NewCafe is Cafe constructor
func NewCafe() *Cafe {
	comps := [8]Computer{}
	ch := make(chan bool, 8)

	for i := 0; i < 8; i++ {
		comps[i] = Computer{nil}
		ch <- true
	}

	return &Cafe{comps, ch}
}

// OccupyComputer simulates a tourist using a computer
func (c *Cafe) OccupyComputer(t *tourist.Tourist) {

	for i := 0; i < 8; i++ {

		// Go to a free computer
		if c.Computers[i].IsFree() {
			fmt.Println("User", t.ID, "occupies the computer", i+1)

			// Set user to t
			c.Computers[i].User = t
			break
		}
	}
}

// KickUser from the computer && make the computer free
func (c *Cafe) KickUser(t *tourist.Tourist) {

	for i := 0; i < 8; i++ {

		// go to a free computer
		if c.Computers[i].User == t {

			fmt.Println("User", t.ID, "is kicked from computer:", i+1)

			// Remove user
			c.Computers[i].User = nil
			break
		}
	}
}
