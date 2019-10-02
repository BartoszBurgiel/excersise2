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
		comps[i] = Computer{nil, true}
		ch <- true
	}

	return &Cafe{comps, ch}
}

// AnyFreeComputer returns an information
// if any of the cafe's computers is free
func (c *Cafe) AnyFreeComputer() (bool, int) {
	for n, com := range c.Computers {
		if com.Free {
			return true, n
		}
	}

	return false, -1
}

// OccupyComputer simulates a tourist using a computer
func (c *Cafe) OccupyComputer(t *tourist.Tourist) {

	for i := 0; i < 8; i++ {

		// go to a free computer
		if c.Computers[i].Free {
			fmt.Println("User", t.ID, "occupies the computer", i+1)

			c.Computers[i].Free = false
			c.Computers[i].User = t
			break
		}
	}
}

func (c *Cafe) KickUser(t *tourist.Tourist) {
	for i := 0; i < 8; i++ {

		// go to a free computer
		if c.Computers[i].User == t {
			fmt.Println("User", t.ID, "is kicked from computer:", i+1)

			c.Computers[i].Free = true
			c.Computers[i].User = nil
			break
		}
	}
}
