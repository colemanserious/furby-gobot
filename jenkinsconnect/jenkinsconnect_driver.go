package jenkinsconnect

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"time"
)

var _ gobot.Driver = (*JenkinsconnectDriver)(nil)

const JobResult string = "jobResult"

type JenkinsconnectDriver struct {
	name       string
	connection gobot.Connection
	halt       chan bool
	gobot.Eventer
	gobot.Commander
}

type JobOutcomeSnapshot struct {
	Outcome JobOutcome
	RanAt   time.Time
}

var jobStates = map[string]JobOutcomeSnapshot{}

func NewJenkinsconnectDriver(a *JenkinsconnectAdaptor, name string) *JenkinsconnectDriver {
	j := &JenkinsconnectDriver{
		name:       name,
		connection: a,
		Eventer:    gobot.NewEventer(),
		Commander:  gobot.NewCommander(),
	}

	j.AddEvent(JobResult)

	j.AddCommand("ParseResults", func(params map[string]interface{}) interface{} {
		var snapshot JobOutcomeSnapshot

		result := ParseJobState(params)
		if (JobOutcome{}) != result {
			fmt.Printf("Result: %v\n", result)
			lastOutcome, ok := jobStates[result.Name]
			if ok {
				fmt.Printf("Last outcome: %v\n", lastOutcome)
			}

			snapshot.Outcome = result
			snapshot.RanAt = time.Now()
			jobStates[result.Name] = snapshot
			gobot.Publish(j.Event(JobResult), result)
		}

		return result
	})

	return j
}

func (j *JenkinsconnectDriver) Name() string { return j.name }

func (j *JenkinsconnectDriver) Connection() gobot.Connection {
	return j.connection
}

func (j *JenkinsconnectDriver) adaptor() *JenkinsconnectAdaptor {
	return j.Connection().(*JenkinsconnectAdaptor)
}

func (j *JenkinsconnectDriver) Start() []error {
	return nil
}

func (j *JenkinsconnectDriver) Halt() []error {
	return nil
}
