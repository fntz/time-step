package tests

import (
	"testing"
	"time"

	timestep "github.com/fntz/time-step/time-step"
	"github.com/stretchr/testify/require"
)

func TestTimeStep(t *testing.T) {
	now := time.Now()
	var steps []time.Time
	tss := timestep.From(now).To(now.Add(3 * time.Hour)).Step(1 * time.Hour)
	for x := range tss.All() {
		steps = append(steps, x)
	}

	require.Len(t, steps, 4)
	require.Equal(t, steps[1].Sub(steps[0]), time.Hour)
	require.Equal(t, steps[2].Sub(steps[1]), time.Hour)
	require.Equal(t, steps[3].Sub(steps[2]), time.Hour)

	// Test Infinity
	max_count := 1000
	i := 0
	tss = timestep.From(now).Infinity().Step(1 * time.Hour)
	for range tss.All() {
		if i > max_count {
			break
		}
		i++
	}
	require.Equal(t, i, max_count+1)
}
