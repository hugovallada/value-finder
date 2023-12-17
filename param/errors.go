package param

type Errors struct {
	Count               uint
	LinesWhereItAppears []uint8
}

func (e *Errors) UpdateCount() *Errors {
	e.Count++
	return e
}

func (e *Errors) UpdateMultipleCounts(counts uint) *Errors {
	e.Count += counts
	return e
}

func (e *Errors) UpdateMultipleLines(lines ...uint8) *Errors {
	e.LinesWhereItAppears = append(e.LinesWhereItAppears, lines...)
	return e
}

func (e *Errors) UpdateLines(line uint8) *Errors {
	e.LinesWhereItAppears = append(e.LinesWhereItAppears, line)
	return e
}

func New(line uint8) Errors {
	return Errors{
		Count:               1,
		LinesWhereItAppears: []uint8{line},
	}
}