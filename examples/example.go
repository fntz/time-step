// Example program demonstrating the time-step package usage
package main

import (
	"fmt"
	"time"

	timestep "github.com/fntz/time-step/time-step"
)

func main() {
	fmt.Println("=== Time-Step Package Examples ===\n")

	// Example 1: Basic finite sequence
	fmt.Println("Example 1: Hourly timestamps for a workday")
	workdayStart := time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC)
	workdayEnd := workdayStart.Add(8 * time.Hour)
	
	fmt.Printf("Workday: %s to %s\n\n", 
		workdayStart.Format("15:04"), 
		workdayEnd.Format("15:04"))
	
	tss := timestep.From(workdayStart).To(workdayEnd).Step(1 * time.Hour)
	
	fmt.Println("Hourly checkpoints:")
	for t := range tss.All() {
		fmt.Printf("  • %s\n", t.Format("15:04"))
	}
	fmt.Println()

	// Example 2: Different time intervals
	fmt.Println("Example 2: Various time intervals")
	
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	
	// Every 15 minutes for 1 hour
	fmt.Println("\nEvery 15 minutes for 1 hour:")
	quarterHour := timestep.From(start).To(start.Add(1 * time.Hour)).Step(15 * time.Minute)
	for t := range quarterHour.All() {
		fmt.Printf("  %s\n", t.Format("15:04"))
	}
	
	// Every 2 days for a week
	fmt.Println("\nEvery 2 days for a week:")
	twoDays := timestep.From(start).To(start.Add(7 * 24 * time.Hour)).Step(48 * time.Hour)
	for t := range twoDays.All() {
		fmt.Printf("  %s\n", t.Format("Jan 02"))
	}
	fmt.Println()

	// Example 3: Infinite sequence with break condition
	fmt.Println("Example 3: Infinite sequence (first 5 values)")
	now := time.Now()
	infinite := timestep.From(now).Infinity().Step(30 * time.Second)
	
	count := 0
	fmt.Println("Next 30-second intervals:")
	for t := range infinite.All() {
		fmt.Printf("  %s\n", t.Format("15:04:05"))
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println()

	// Example 4: Collecting results
	fmt.Println("Example 4: Collecting timestamps into a slice")
	
	meetingStart := time.Date(2024, 1, 15, 14, 0, 0, 0, time.UTC)
	meetingEnd := meetingStart.Add(1 * time.Hour)
	
	meetingTimes := timestep.From(meetingStart).To(meetingEnd).Step(15 * time.Minute)
	
	var agendaTimes []time.Time
	for t := range meetingTimes.All() {
		agendaTimes = append(agendaTimes, t)
	}
	
	fmt.Printf("Meeting agenda times (%d items):\n", len(agendaTimes))
	for i, t := range agendaTimes {
		fmt.Printf("  Agenda item %d: %s\n", i+1, t.Format("15:04"))
	}
	fmt.Println()

	// Example 5: Real-world use case - generating report schedule
	fmt.Println("Example 5: Generating weekly report schedule")
	
	today := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC) // Monday
	nextMonth := today.Add(30 * 24 * time.Hour)
	
	fmt.Println("Weekly reports (Mondays at 9:00 AM):")
	weeklyReports := timestep.From(today).To(nextMonth).Step(7 * 24 * time.Hour)
	
	reportCount := 0
	for t := range weeklyReports.All() {
		fmt.Printf("  Week %d: %s\n", reportCount+1, t.Format("Jan 02, 2006"))
		reportCount++
	}
	
	fmt.Printf("\nTotal reports scheduled: %d\n", reportCount)
	fmt.Println("\n=== End of Examples ===")
}
