package jenkinsconnect

import (
	"fmt"
	//	"log"
)

type JobOutcome struct {
	State JobState
	Name  string
}

type JobState int

const (
	UNKNOWN JobState = iota
	STARTED
	COMPLETED
	FINALIZED
	FAILED
	SUCCESS
	ABORTED
	UNSTABLE
)

var jobStateValues = []string{
	UNKNOWN:   "UNKNOWN",
	STARTED:   "STARTED",
	COMPLETED: "COMPLETED",
	FINALIZED: "FINALIZED",
	FAILED:    "FAILURE",
	SUCCESS:   "SUCCESS",
	ABORTED:   "ABORTED",
	UNSTABLE:  "UNSTABLE",
}

func (j JobState) String() string {
	if j <= 0 || int(j) >= len(jobStateValues) {
		return "unknown"
	}
	return jobStateValues[j]
}

func getJobState(s string) JobState {

	for index, _ := range jobStateValues {
		if s == jobStateValues[index] {
			return JobState(index)
		}
	}

	return UNKNOWN
}

func findKey(key string, value interface{}, checkKey string) (string, bool) {
	fmt.Printf("findKey for %v against checkKey %v\n", key, checkKey)
	switch val := value.(type) {
	case string:
		return val, (key == checkKey)
	case map[string]interface{}:
		for newKey, newValue := range val {
			foundValue, found := findKey(newKey, newValue, checkKey)
			if found {
				return foundValue, found
			}
		}
	}
	return "", false
}

func ParseJobState(params map[string]interface{}) JobOutcome {

	// Example: non-compiling code: [1]JobOutcome != []JobOutcome.
	// Size of array is included in its type.  Use slices instead
	//outcomesArray := [...]JobOutcome{firstOutcome}
	//return outcomesArray

	fmt.Printf("Finding JobState... %v\n", params)
	var statusFound, nameFound, phaseFound bool
	var statusValue, nameValue, phaseValue string

	var outcome, blankOutcome JobOutcome

	for key, value := range params {
		if !phaseFound {
			phaseValue, phaseFound = findKey(key, value, "phase")
			if phaseFound {
				// we get 3 entries, with phases STARTED, COMPLETED, and FINALIZED.
				//  Only the FINALIZED one will give us real info
				if phaseValue != "FINALIZED" {
					return blankOutcome
				}
			}
		}
		if !statusFound {
			statusValue, statusFound = findKey(key, value, "status")
		}
		if !nameFound {
			nameValue, nameFound = findKey(key, value, "name")
		}
		if nameFound && statusFound {
			fmt.Printf("name: %v, status: %v\n", nameValue, statusValue)

			outcome.State = getJobState(statusValue)
			outcome.Name = nameValue
			break
		}
	}

	return outcome
}
