package audio

import (
	"github.com/hybridgroup/gobot"
	"time"
)

var _ gobot.Driver = (*AudioDriver)(nil)

type AudioDriver struct {
	name       string
	connection gobot.Connection
	interval   time.Duration
	halt       chan bool
	gobot.Eventer
	gobot.Commander
}

func NewAudioDriver(a *AudioAdaptor, name string) *AudioDriver {
	d := &AudioDriver{
		name:       name,
		connection: a,
		interval:   500 * time.Millisecond,
		halt:       make(chan bool, 0),
		Eventer:    gobot.NewEventer(),
		Commander:  gobot.NewCommander(),
	}

	return d
}

func (d *AudioDriver) Name() string { return d.name }

func (d *AudioDriver) Connection() gobot.Connection {
	return d.connection
}

func (d *AudioDriver) Sound(fileName string) []error {
	return d.Connection().(*AudioAdaptor).Sound(fileName)
}

func (d *AudioDriver) adaptor() *AudioAdaptor {
	return d.Connection().(*AudioAdaptor)
}

func (d *AudioDriver) Start() (err []error) {
	return
}

func (d *AudioDriver) Halt() (err []error) {
	return
}
