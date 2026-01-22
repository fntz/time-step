package timestep

import "time"

type startTimeStep struct {
	start time.Time
}

func From(start time.Time) *startTimeStep {
	return &startTimeStep{
		start: start,
	}
}

func (s *startTimeStep) To(end time.Time) *startEndTimeStep {
	return &startEndTimeStep{
		start: s.start,
		end:   &end,
	}
}

func (s *startTimeStep) Infinity() *startEndTimeStep {
	return &startEndTimeStep{
		start: s.start,
	}
}

type startEndTimeStep struct {
	start time.Time
	end   *time.Time
}

func (s *startEndTimeStep) Step(step time.Duration) *TimeStep {
	return &TimeStep{
		start: s.start,
		end:   s.end,
		step:  step,
	}
}
