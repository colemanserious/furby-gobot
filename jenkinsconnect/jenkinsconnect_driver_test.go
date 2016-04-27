package jenkinsconnect

import (
	"testing"
	"time"

	"github.com/hybridgroup/gobot"
)

func TestJenkinsconnectDriver(t *testing.T) {
	d := NewJenkinsconnectDriver(NewJenkinsconnectAdaptor("conn"), "dev")

	gobot.Assert(t, d.Name(), "dev")
	gobot.Assert(t, d.Connection().Name(), "conn")

	gobot.Assert(t, len(d.Start()), 0)

	<-time.After(d.interval)

	sem := make(chan bool, 0)

	gobot.On(d.Event(Hello), func(data interface{}) {
		sem <- true
	})

	select {
	case <-sem:
	case <-time.After(600 * time.Millisecond):
		t.Errorf("Hello Event was not published")
	}

	gobot.Assert(t, len(d.Halt()), 0)

	gobot.On(d.Event(Hello), func(data interface{}) {
		sem <- true
	})

	select {
	case <-sem:
		t.Errorf("Hello Event should not publish after Halt")
	case <-time.After(600 * time.Millisecond):
	}
}
