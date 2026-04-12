// types to use in this program.

package store

import (
	"time"
)

type Process struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	CreatedBy   int // Tied to the ID value in the User table
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
	CompletedBy int // user name
	CompletedAt time.Time
	Notes       string // optional notes from operator
}

type User struct {
	ID       int
	Username string
	Role     string //admin or operator
	Name     string
}

type WorkOrder struct {
	ID         int
	ProcessID  int    // The template this job is following
	Identifier string // e.g., "Truck SN-987" or "Customer Job #442"
	Status     string // "Not Started", "In Progress", "Completed"
	CreatedAt  time.Time
	CreatedBy  int // User ID of whoever issued the work order
}

type QAPhoto struct {
	ID           int
	CompletionID int    // Ties the photo directly to the step sign-off
	FilePath     string // e.g., "/uploads/qa/workorder_12/step_4_xyz.jpg"
	UploadedAt   time.Time
	UploadedBy   int // User ID
}

type Asset struct {
	ID           int
	SerialNumber string
	CustomerName string
	DateShipped  time.Time
	// This allows you to tie multiple WorkOrders to a single asset over its lifespan
}
