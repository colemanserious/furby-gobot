package jenkins

import ()

type JobOutcome struct {
	State JobState
	Name  string
}

type JobState int

const (
	_                      = iota
	STARTED       JobState = 0
	COMPLETED     JobState = 1
	FINALIZED     JobState = 2
	FAILED        JobState = 3
	PASSING       JobState = 4
	STILL_FAILING JobState = 14
)

var (
	jobStateMachine map[string]JobOutcome
)

func init() {
	jobStateMachine = make(map[string]JobOutcome)
}

func (j JobState) String() string {
	switch {
	case j == FAILED:
		return "FAILED"
	case j == FINALIZED:
		return "FINALIZED"
	case j == COMPLETED:
		return "COMPLETED"
	case j == STARTED:
		return "STARTED"
	case j == PASSING:
		return "PASSING"
	case j == STILL_FAILING:
		return "BONEHEAD FAILING"
	}
	return "UNKNOWN!"
}

func ParseJobState(params map[string]interface{}) []JobOutcome {
	// iterate over map, get final result for each job.
	var firstOutcome JobOutcome
	firstOutcome.State = COMPLETED
	firstOutcome.Name = "Test job"

	// Example: non-compiling code: [1]JobOutcome != []JobOutcome.
	// Size of array is included in its type.  Use slices instead
	//outcomesArray := [...]JobOutcome{firstOutcome}
	//return outcomesArray

	jobStateMachine[firstOutcome.Name] = firstOutcome
	outcomesSlice := make([]JobOutcome, 1)
	outcomesSlice[0] = firstOutcome
	return outcomesSlice
}
