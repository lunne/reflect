package reflect

import "time"

// RecurrentGoal represents a goal that is repeated until a certain date.
// Every daily check-in creates new goals for the user based of RecurrentGoals.
// When creating a new goal from a RecurrentGoal the DueDate on the new goal is set to
// today+TimeToComplete.
type RecurrentGoal struct {
	Id             string
	Description    string
	Until          time.Time
	TimeToComplete time.Duration
}

// Goal represents a goal that a user has set for themselves.
// A goal can be recurrent, in which case it will be repeated until a certain date.
// A goal can also be completed, in which case it will be marked as completed.
// To complete a goal a user must create a report and mark the goal as completed.
type Goal struct {
	Id          string
	Description string
	DueDate     time.Time
	Completed   time.Time
	ReportId    string
}

// Report is the report that a user generates for a given day that is tied to one or more goals.
type Report struct {
	Id string

	Description string
	Comment     string

	// Add fields to match the questions in readme
	Goals     []Goal
	CreatedAt time.Time
}
