package timestep

import (
	"iter"
	"time"
)

type TimeStep struct {
	start time.Time
	end   *time.Time
	step  time.Duration
}

func (s *TimeStep) All() iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		current := s.start
		for {
			if s.end != nil && current.After(*s.end) {
				return
			}
			if !yield(current) {
				return
			}
			current = current.Add(s.step)
		}
	}
}
