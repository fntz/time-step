# Time-Step

A Go package providing a fluent DSL for generating time sequences with configurable start, end, and step intervals using Go 1.24's new iterator pattern.

## Features

- **Fluent API**: Chainable methods for intuitive time sequence generation
- **Iterator Support**: Uses Go 1.24's `iter.Seq` for efficient iteration
- **Flexible Ranges**: Support for both finite and infinite time sequences
- **Type Safety**: Built with Go's native `time.Time` and `time.Duration` types
- **Zero Dependencies**: Pure Go implementation with minimal external dependencies

## Installation

```bash
go get github.com/fntz/time-step
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"
    "github.com/fntz/time-step/time-step"
)

func main() {
    now := time.Now()
    
    // Generate hourly timestamps for the next 3 hours
    tss := timestep.From(now).To(now.Add(3 * time.Hour)).Step(1 * time.Hour)
    
    for t := range tss.All() {
        fmt.Println(t.Format("2006-01-02 15:04:05"))
    }
}
```

## API Reference

### Core Types

#### `TimeStep`
The main struct that represents a time sequence configuration.

```go
type TimeStep struct {
    start time.Time
    end   *time.Time  // nil for infinite sequences
    step  time.Duration
}
```

#### Methods

##### `All() iter.Seq[time.Time]`
Returns an iterator that yields time values according to the configured sequence.

### DSL Methods

#### `From(start time.Time) *startTimeStep`
Creates a new time sequence starting at the specified time.

#### `(s *startTimeStep) To(end time.Time) *startEndTimeStep`
Sets the end time for a finite sequence.

#### `(s *startTimeStep) Infinity() *startEndTimeStep`
Creates an infinite sequence (no end time).

#### `(s *startEndTimeStep) Step(step time.Duration) *TimeStep`
Sets the step duration between consecutive time values.

## Usage Examples

### Basic Finite Sequence

```go
start := time.Now()
end := start.Add(6 * time.Hour)

// Generate every 2 hours from start to end
tss := timestep.From(start).To(end).Step(2 * time.Hour)

for t := range tss.All() {
    fmt.Println(t)
}
// Output:
// 2026-01-22 15:36:27.143476276 +0300 +03 m=+0.000032617
// 2026-01-22 17:36:27.143476276 +0300 +03 m=+7200.000032617
// 2026-01-22 19:36:27.143476276 +0300 +03 m=+14400.000032617
// 2026-01-22 21:36:27.143476276 +0300 +03 m=+21600.000032617

```

### Infinite Sequence

```go
start := time.Now()

// Generate timestamps every 5 minutes indefinitely
tss := timestep.From(start).Infinity().Step(5 * time.Minute)

count := 0
for t := range tss.All() {
    fmt.Println(t.Format("15:04:05"))
    count++
    if count >= 10 {
        break // Important: break infinite loops!
    }
}

// output: 
// 15:37:05
// 15:42:05
// 15:47:05
// 15:52:05
// 15:57:05
```

## Requirements

- Go 1.24 or later (requires `iter` package support)
- No external dependencies beyond standard library

## Support

For issues, questions, or feature requests, please open an issue on the GitHub repository.

## License

MIT License
