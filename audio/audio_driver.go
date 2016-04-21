package audio

import (
	"github.com/hybridgroup/gobot"
	"time"
)

var _ gobot.Driver = (*AudioDriver)(nil)

const Hello string = "hello"

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

	d.AddEvent(Hello)

	d.AddCommand(Hello, func(params map[string]interface{}) interface{} {
		return d.Hello()
	})

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

func (d *AudioDriver) Hello() string {
	return "hello from " + d.Name() + "!"
}

func (d *AudioDriver) Start() []error {
	go func() {
		for {
			gobot.Publish(d.Event(Hello), d.Hello())

			select {
			case <-time.After(d.interval):
			case <-d.halt:
				return
			}
		}
	}()
	return nil
}

func (d *AudioDriver) Halt() []error {
	d.halt <- true
	return nil
}
