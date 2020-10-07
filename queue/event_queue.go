package queue

import "time"

var LogLines = Statistics{}

type Statistics struct {
	LastScan time.Time
	Imports  []Import
}

type Import struct {
	Text string
}

func (s *Statistics) AddImport(i Import) {
	const n = 250
	if len(s.Imports) >= n {
		copy(s.Imports, s.Imports[len(s.Imports)-n+1:])
		s.Imports = s.Imports[:n-1]
	}
	s.Imports = append(s.Imports, i)
}
