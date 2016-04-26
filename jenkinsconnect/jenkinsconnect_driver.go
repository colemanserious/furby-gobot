package jenkinsconnect

import (
	//"fmt"
	"log"
	"time"

	"github.com/hybridgroup/gobot"
)

var _ gobot.Driver = (*JenkinsconnectDriver)(nil)

const Hello string = "hello"

type JenkinsconnectDriver struct {
	name       string
	connection gobot.Connection
	interval   time.Duration
	halt       chan bool
	gobot.Eventer
	gobot.Commander
}

func NewJenkinsconnectDriver(a *JenkinsconnectAdaptor, name string) *JenkinsconnectDriver {
	j := &JenkinsconnectDriver{
		name:       name,
		connection: a,
		interval:   500 * time.Millisecond,
		halt:       make(chan bool, 0),
		Eventer:    gobot.NewEventer(),
		Commander:  gobot.NewCommander(),
	}

	j.AddEvent(Hello)

	j.AddCommand("ParseResults", func(params map[string]interface{}) interface{} {
		log.Println("Received result from Jenkins.")
		//log.Printf("Received result from Jenkins... %v", params)
		result := ParseJobState(params)
		//log.Printf("Result received: %v", result[0].State)
		return result
		//return fmt.Sprintf("Parsed results:  %v", result[0])
		//return fmt.Sprintf("Parsing Jenkins results: %v", jenkins.Name)
	})

	j.AddCommand(Hello, func(params map[string]interface{}) interface{} {
		//	return j.Hello()
		return nil
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
	go func() {
		for {
			//gobot.Publish(j.Event(Hello), j.Hello())

			select {
			case <-time.After(j.interval):
			case <-j.halt:
				return
			}
		}
	}()
	return nil
}

func (j *JenkinsconnectDriver) Halt() []error {
	j.halt <- true
	return nil
}
