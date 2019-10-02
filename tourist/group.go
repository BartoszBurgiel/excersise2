package tourist

// Group represents a group of tourists
type Group struct {
	Tourists [25]*Tourist

	//Every tourist used a computer
	IsDone bool

	TouristsToGo, TouristsDone int
}
