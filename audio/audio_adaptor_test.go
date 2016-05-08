package audio

import (
	"os/exec"
	"testing"

	"github.com/hybridgroup/gobot"
)

func TestAudioAdaptor(t *testing.T) {
	a := NewAudioAdaptor("tester")

	gobot.Assert(t, a.Name(), "tester")

	gobot.Assert(t, len(a.Connect()), 0)

	_, err := exec.LookPath("aplay")
	numErrsForTest := 0
	if err != nil {
		numErrsForTest = 1
	}
	gobot.Assert(t, len(a.Sound("../resources/foo.wav")), numErrsForTest)

	gobot.Assert(t, len(a.Connect()), 0)

	gobot.Assert(t, len(a.Finalize()), 0)
}
