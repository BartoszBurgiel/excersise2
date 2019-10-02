package tourist

// Tourist struct defines a tourist
type Tourist struct {
	ID int

	// TimeOnline -> minutes already spent online
	// LimitOnline -> max time spent online
	TimeOnline, LimitOnline int
}
