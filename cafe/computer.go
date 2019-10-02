package cafe

import "internetCafe/tourist"

// Computer struct defines a computer from the internet cafe
type Computer struct {
	User *tourist.Tourist
	Free chan bool
}

// NewUser sets the attributes of Computer to a new user
func (c *Computer) NewUser(t *tourist.Tourist) {
	c.Free <- false
	c.User = t
}
