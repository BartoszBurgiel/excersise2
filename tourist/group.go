package tourist

// Group represents a group of tourists
type Group struct {
	Tourists [25]*Tourist

	//Every tourist used a computer
	IsDone bool

	TouristsToGo, TouristsDone int
}

// NewGroup constructor
func NewGroup() *Group {
	tourists := [25]*Tourist{}

	for i := 0; i < 25; i++ {
		tourists[i] = &Tourist{i + 1, 0, false, false, true}
	}

	return &Group{tourists, false, 25, 0}
}
