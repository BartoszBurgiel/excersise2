package cafe

// Cafe represents an internetcafe with its computers
type Cafe struct {
	Computers    [8]Computer
	FreeComputer int
}

// AnyComputerFree checks if any computer in the Cafe is free
func (c Cafe) AnyComputerFree() bool {
	for _, comp := range c.Computers {
		if <-comp.Free {
			return true
		}
	}

	return false
}
