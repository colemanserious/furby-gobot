package jenkinsconnect

import (
	"testing"

	"github.com/hybridgroup/gobot"
)

func TestJenkinsconnectAdaptor(t *testing.T) {
	a := NewJenkinsconnectAdaptor("tester")

	gobot.Assert(t, a.Name(), "tester")

	gobot.Assert(t, len(a.Connect()), 0)

	gobot.Assert(t, a.Ping(), "pong")

	gobot.Assert(t, len(a.Connect()), 0)

	gobot.Assert(t, len(a.Finalize()), 0)
}
