package jenkinsconnect

import (
	"testing"

	"github.com/hybridgroup/gobot"
)

func TestJenkinsconnectDriver(t *testing.T) {
	d := NewJenkinsconnectDriver(NewJenkinsconnectAdaptor("conn"), "dev")

	gobot.Assert(t, d.Name(), "dev")
	gobot.Assert(t, d.Connection().Name(), "conn")

	gobot.Assert(t, len(d.Start()), 0)

	gobot.Assert(t, len(d.Halt()), 0)

}
