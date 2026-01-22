package timestep

import (
	"testing"
	"time"
)

func BenchmarkTimeStep_Finite(b *testing.B) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour) // 24 hours total
	step := time.Hour                 // 25 iterations (including start and end)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tss := From(start).To(end).Step(step)
		count := 0
		for range tss.All() {
			count++
		}
		// Verify we got the expected number of iterations
		if count != 25 {
			b.Errorf("expected 25 iterations, got %d", count)
		}
	}
}

func BenchmarkTimeStep_Infinite(b *testing.B) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	step := time.Second
	const iterations = 1000 // Fixed number of iterations for benchmarking

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tss := From(start).Infinity().Step(step)
		count := 0
		for range tss.All() {
			count++
			if count >= iterations {
				break
			}
		}
		if count != iterations {
			b.Errorf("expected %d iterations, got %d", iterations, count)
		}
	}
}

func BenchmarkTimeStep_SmallStep(b *testing.B) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(time.Hour)
	step := time.Millisecond // 3,600,000 iterations

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tss := From(start).To(end).Step(step)
		count := 0
		// Limit iterations to avoid excessive runtime
		for range tss.All() {
			count++
			if count >= 10000 {
				break
			}
		}
	}
}

func BenchmarkTimeStep_LargeRange(b *testing.B) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(365 * 24 * time.Hour) // One year
	step := 24 * time.Hour                 // Daily for a year (366 days including leap year)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tss := From(start).To(end).Step(step)
		count := 0
		for range tss.All() {
			count++
		}
		// 2024 is a leap year, so 366 days
		if count != 366 {
			b.Errorf("expected 366 iterations, got %d", count)
		}
	}
}

func BenchmarkTimeStep_Construction(b *testing.B) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)
	step := time.Hour

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = From(start).To(end).Step(step)
	}
}

func BenchmarkTimeStep_IterationOnly(b *testing.B) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)
	step := time.Hour
	tss := From(start).To(end).Step(step)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		for range tss.All() {
			count++
		}
		if count != 25 {
			b.Errorf("expected 25 iterations, got %d", count)
		}
	}
}