package jenkinsconnect

import (
	"fmt"
	//	"log"
	"strconv"
)

type JobOutcome struct {
	State JobState
	Name  string
}

type JobState int

const (
	_                  = iota
	STARTED   JobState = 0
	COMPLETED JobState = 1
	FINALIZED JobState = 2
	FAILED    JobState = 3
	PASSING   JobState = 4
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
	}
	return "UNKNOWN!"
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

	fmt.Println("Finding JobState...")
	var statusFound, nameFound bool
	var statusValue, nameValue string

	var outcome JobOutcome

	for key, value := range params {
		if !statusFound {
			statusValue, statusFound = findKey(key, value, "status")
		}
		if !nameFound {
			nameValue, nameFound = findKey(key, value, "name")
		}
		if nameFound && statusFound {
			fmt.Printf("name: %v, status: %v\n", nameValue, statusValue)

			// todo: This is converting "1" to 1, not FAILED to 3
			statusInt, _ := strconv.Atoi(statusValue)
			outcome.State = JobState(statusInt)
			outcome.Name = nameValue
			break
		}
	}

	return outcome
}
