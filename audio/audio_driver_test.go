package audio

import (
	"os/exec"
	"testing"
	"time"

	"github.com/hybridgroup/gobot"
)

func TestAudioDriver(t *testing.T) {
	d := NewAudioDriver(NewAudioAdaptor("conn"), "dev")

	gobot.Assert(t, d.Name(), "dev")
	gobot.Assert(t, d.Connection().Name(), "conn")

	ret := d.Command(Hello)(nil)
	gobot.Assert(t, ret.(string), "hello from dev!")

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

	_, err := exec.LookPath("aplayer")
	numErrsForTest := 0
	if err != nil {
		numErrsForTest = 1
	}
	gobot.Assert(t, len(d.Sound("foo.wav")), numErrsForTest)
}
