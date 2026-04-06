// types to use in this program.

package main

import "time"

type Process struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	CreatedBy   string
	Steps       []Step
}

type Step struct {
	ID          int
	ProcessID   int // which process it belongs to
	Name        string
	Description string
	Required    bool
	Critical    bool // if missed = termination level
	Order       int  // sequence in the process
}

type Completion struct {
	ID          int
	ProcessID   int
	StepID      int
	CompletedBy string // user name
	CompletedAt time.Time
	Notes       string // optional notes from operator
}

type User struct {
	ID       int
	Username string
	Role     string //admin or operator
	Name     string
}
