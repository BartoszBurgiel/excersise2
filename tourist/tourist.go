package tourist

// Tourist struct defines a tourist
type Tourist struct {
	ID int

	// TimeOnline -> minutes
	TimeOnline int

	WasOnline, IsOnline, Waits bool
}

// Status messages

//SetAsDone updates tourist's after he is done with the computer
func (t *Tourist) SetAsDone() {
	t.WasOnline = true
	t.IsOnline = false
	t.Waits = false
}
