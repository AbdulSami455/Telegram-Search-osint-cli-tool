package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ProgressBar represents a visual progress bar
type ProgressBar struct {
	total     int
	current   int
	width     int
	prefix    string
	suffix    string
	startTime time.Time
}

// NewProgressBar creates a new progress bar
func NewProgressBar(total int, prefix string) *ProgressBar {
	return &ProgressBar{
		total:     total,
		current:   0,
		width:     50,
		prefix:    prefix,
		suffix:    "",
		startTime: time.Now(),
	}
}

// Update updates the progress bar
func (pb *ProgressBar) Update(current int) {
	pb.current = current
	pb.render()
}

// Increment increments the progress by 1
func (pb *ProgressBar) Increment() {
	pb.current++
	pb.render()
}

// render renders the progress bar to stdout
func (pb *ProgressBar) render() {
	percentage := float64(pb.current) / float64(pb.total)
	filled := int(float64(pb.width) * percentage)

	bar := strings.Repeat("█", filled) + strings.Repeat("░", pb.width-filled)

	elapsed := time.Since(pb.startTime)
	eta := time.Duration(0)
	if pb.current > 0 {
		eta = time.Duration(float64(elapsed) * float64(pb.total-pb.current) / float64(pb.current))
	}

	fmt.Fprintf(os.Stdout, "\r%s [%s] %d/%d (%.1f%%) ETA: %s%s",
		pb.prefix, bar, pb.current, pb.total, percentage*100, eta.Round(time.Second), pb.suffix)

	if pb.current >= pb.total {
		fmt.Println() // New line when complete
	}
}

// Spinner represents a loading spinner
type Spinner struct {
	frames []string
	pos    int
	prefix string
	active bool
	stop   chan bool
}

// NewSpinner creates a new spinner
func NewSpinner(prefix string) *Spinner {
	return &Spinner{
		frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		pos:    0,
		prefix: prefix,
		active: false,
		stop:   make(chan bool),
	}
}

// Start starts the spinner
func (s *Spinner) Start() {
	if s.active {
		return
	}
	s.active = true

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(os.Stdout, "\r%s %s", s.prefix, s.frames[s.pos])
				s.pos = (s.pos + 1) % len(s.frames)
			case <-s.stop:
				return
			}
		}
	}()
}

// Stop stops the spinner
func (s *Spinner) Stop() {
	if !s.active {
		return
	}
	s.active = false
	s.stop <- true
	fmt.Fprintf(os.Stdout, "\r%s ✓\n", s.prefix)
}

// Status represents a status message
type Status struct {
	message string
	level   string
}

// NewStatus creates a new status message
func NewStatus(message, level string) *Status {
	return &Status{
		message: message,
		level:   level,
	}
}

// Print prints the status message with appropriate formatting
func (s *Status) Print() {
	switch s.level {
	case "info":
		fmt.Printf("ℹ️  %s\n", s.message)
	case "success":
		fmt.Printf("✅ %s\n", s.message)
	case "warning":
		fmt.Printf("⚠️  %s\n", s.message)
	case "error":
		fmt.Printf("❌ %s\n", s.message)
	default:
		fmt.Printf("%s\n", s.message)
	}
}

// PrintInfo prints an info message
func PrintInfo(message string) {
	NewStatus(message, "info").Print()
}

// PrintSuccess prints a success message
func PrintSuccess(message string) {
	NewStatus(message, "success").Print()
}

// PrintWarning prints a warning message
func PrintWarning(message string) {
	NewStatus(message, "warning").Print()
}

// PrintError prints an error message
func PrintError(message string) {
	NewStatus(message, "error").Print()
}

// PrintStep prints a step in a multi-step process
func PrintStep(step int, total int, message string) {
	fmt.Printf("Step %d/%d: %s\n", step, total, message)
}

// PrintHeader prints a section header
func PrintHeader(title string) {
	fmt.Printf("\n%s\n", strings.Repeat("=", len(title)+4))
	fmt.Printf("= %s =\n", title)
	fmt.Printf("%s\n\n", strings.Repeat("=", len(title)+4))
}

// PrintSubHeader prints a subsection header
func PrintSubHeader(title string) {
	fmt.Printf("\n%s\n", strings.Repeat("-", len(title)+2))
	fmt.Printf("- %s -\n", title)
	fmt.Printf("%s\n\n", strings.Repeat("-", len(title)+2))
}

// ClearLine clears the current line
func ClearLine() {
	fmt.Fprintf(os.Stdout, "\r%s\r", strings.Repeat(" ", 100))
}
