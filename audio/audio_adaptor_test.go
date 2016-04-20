package audio

import (
	"testing"

	"github.com/hybridgroup/gobot"
)

func TestAudioAdaptor(t *testing.T) {
	a := NewAudioAdaptor("tester")

	gobot.Assert(t, a.Name(), "tester")

	gobot.Assert(t, len(a.Connect()), 0)

	gobot.Assert(t, len(a.Sound("foo.wav")), 0)

	gobot.Assert(t, len(a.Connect()), 0)

	gobot.Assert(t, len(a.Finalize()), 0)
}
