package tourist

// Tourist struct defines a tourist
type Tourist struct {
	// 'Name' of the tourist
	ID int

	// TimeOnline -> minutes already spent online
	// LimitOnline -> max time spent online
	TimeOnline, LimitOnline int
}
