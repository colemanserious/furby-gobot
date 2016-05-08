package audio

import (
	"os/exec"
	"testing"

	"github.com/hybridgroup/gobot"
)

func TestAudioDriver(t *testing.T) {
	d := NewAudioDriver(NewAudioAdaptor("conn"), "dev", nil)

	gobot.Assert(t, d.Name(), "dev")
	gobot.Assert(t, d.Connection().Name(), "conn")

	gobot.Assert(t, len(d.Start()), 0)

	gobot.Assert(t, len(d.Halt()), 0)

	_, err := exec.LookPath("aplay")
	numErrsForTest := 0
	if err != nil {
		numErrsForTest = 1
	}
	gobot.Assert(t, len(d.Sound("../resources/foo.wav")), numErrsForTest)
}
