package cafe

import (
	"fmt"
	"internetCafe/tourist"
)

// Cafe represents an internetcafe with its computers
type Cafe struct {
	Computers    []Computer
	FreeComputer chan bool
}

// NewCafe is Cafe constructor
func NewCafe(nComp int) *Cafe {
	comps := []Computer{}
	ch := make(chan bool, nComp)

	for i := 0; i < nComp; i++ {
		comps = append(comps, Computer{nil})
		ch <- true
	}

	return &Cafe{comps, ch}
}

// OccupyComputer simulates a tourist using a computer
func (c *Cafe) OccupyComputer(t *tourist.Tourist) {

	for i := 0; i < len(c.Computers); i++ {

		// Go to a free computer
		if c.Computers[i].IsFree() {
			fmt.Printf("User %d occupies the computer nr. %d now\n\n", t.ID, i+1)

			// Set user to t
			c.Computers[i].User = t
			break
		}
	}
}

// KickUser from the computer and make the computer free
func (c *Cafe) KickUser(t *tourist.Tourist) {

	for i := 0; i < len(c.Computers); i++ {

		// go to a free computer
		if c.Computers[i].User == t {

			fmt.Printf(" > User %d is kicked from the computer nr. %d\n   He spent %d minutes online\n\n", t.ID, i+1, t.TimeOnline)

			// Remove user
			c.Computers[i].User = nil
			break
		}
	}
}

// GetUser returns a pointer to the user using computer with the index n
func (c *Cafe) GetUser(n int) *tourist.Tourist {
	return c.Computers[n].User
}
