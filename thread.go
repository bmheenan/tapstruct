package taps

// There are three ways to hold a thread's data

// Threadrow holds the info for a thread when it needs to be displayed in a hierarchical table with other threads
type Threadrow struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	State    ThreadState `json:"state"`
	Cost     int         `json:"costCtx"` // Different for each stakeholder + total; depends on the context
	Owner    Stakeholder `json:"owner"`
	Iter     string      `json:"iter"` // The base iteration of the thread, not where its appearing
	Ord      int         `json:"ord"`
	Children []Threadrow `json:"children"`
}

// Threaddetail holds detailed information on a thread, needed for its thread view
type Threaddetail struct {
	ID           int64                  `json:"id"`
	Name         string                 `json:"name"`
	Desc         string                 `json:"desc"`
	State        ThreadState            `json:"state"`
	CostDir      int                    `json:"costDirect"`
	CostTot      int                    `json:"costTotal"` // Includes all descendants
	Owner        Stakeholder            `json:"owner"`
	Stakeholders map[string]Stakeholder `json:"stakeholders"`
	Iter         string                 `json:"iteration"`
	Percentile   float32                `json:"percentile"`
}

// Threadrel tracks the thread data the backend uses for calculating cost, order, and percentile
type Threadrel struct {
	ID           int64
	State        string
	CostDir      int
	Owner        string
	Iter         string
	Percentile   float64
	Stakeholders map[string](struct {
		Iter string
		Ord  int
	})
	Parents map[int64](struct {
		Iter string
		Ord  int
	})
}

// ThreadState describes the possible states for a thread
type ThreadState string

const (
	// NotStarted is the default state; it hasn't been worked on (yet)
	NotStarted ThreadState = "not started"
	// InProgress is for when a thread is being worked on, but is not complete
	InProgress ThreadState = "in progress"
	// Done is for when the value has been delivered to the stakeholders
	Done ThreadState = "done"
	// Closed is appropriate for threads working as intended, unactionable, or duplicates
	Closed ThreadState = "closed"
	// Archived is for threads that were valid, but low enough priority that they were never addressed
	Archived ThreadState = "archived"
)
