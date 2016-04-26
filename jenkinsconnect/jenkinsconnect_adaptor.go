package jenkinsconnect

import (
	"github.com/hybridgroup/gobot"
)

var _ gobot.Adaptor = (*JenkinsconnectAdaptor)(nil)

type JenkinsconnectAdaptor struct {
	name string
}

func NewJenkinsconnectAdaptor(name string) *JenkinsconnectAdaptor {
	return &JenkinsconnectAdaptor{
		name: name,
	}
}

func (j *JenkinsconnectAdaptor) Name() string { return j.name }

func (j *JenkinsconnectAdaptor) Connect() []error {
	return nil
}

func (j *JenkinsconnectAdaptor) Finalize() []error { return nil }

func (j *JenkinsconnectAdaptor) Ping() string { return "pong" }
