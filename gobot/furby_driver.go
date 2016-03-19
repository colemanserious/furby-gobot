package furby

import (
	"time"

	"github.com/hybridgroup/gobot"
)

var _ gobot.Driver = (*FurbyDriver)(nil)

const Hello string = "hello"

type FurbyDriver struct {
	name string
	connection gobot.Connection
	interval time.Duration
	halt chan bool
	gobot.Eventer
	gobot.Commander
}

func NewFurbyDriver(a *FurbyAdaptor, name string) *FurbyDriver {
	f := &FurbyDriver{
		name: name,
		connection: a,
		interval: 500*time.Millisecond,
		halt: make(chan bool, 0),
    Eventer:    gobot.NewEventer(),
    Commander:  gobot.NewCommander(),
	}

	f.AddEvent(Hello)

	f.AddCommand(Hello, func(params map[string]interface{}) interface{} {
		return f.Hello()
	})

	return f
}

func (f *FurbyDriver) Name() string { return f.name }

func (f *FurbyDriver) Connection() gobot.Connection {
	return f.connection
}

func (f *FurbyDriver) adaptor() *FurbyAdaptor {
	return f.Connection().(*FurbyAdaptor)
}

func (f *FurbyDriver) Hello() string {
	return "hello from " + f.Name() + "!"
}

func (f *FurbyDriver) Ping() string {
	return f.adaptor().Ping()
}

func (f *FurbyDriver) Start() []error {
	go func() {
		for {
			gobot.Publish(f.Event(Hello), f.Hello())

			select {
			case <- time.After(f.interval):
			case <- f.halt:
				return
			}
		}
	}()
	return nil
}

func (f *FurbyDriver) Halt() []error {
	f.halt <- true
	return nil
}

