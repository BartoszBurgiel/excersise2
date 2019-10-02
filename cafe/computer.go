package cafe

import "internetCafe/tourist"

// Computer struct defines a computer from the internet cafe
type Computer struct {

	// User occupying the computer
	User *tourist.Tourist
}

// IsFree tells if a computer is currently used by a tourist
func (c *Computer) IsFree() bool {
	if c.User != nil {
		return false
	}

	return true
}
