package furby

import (
	"errors"
	"testing"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func initTestFurbyDriver(conn gpio.DigitalWriter, queue chan string) *FurbyDriver {
	testAdaptorDigitalWrite = func() (err error) {
		return nil
	}
	testAdaptorPwmWrite = func() (err error) {
		return nil
	}
	return NewFurbyDriver(conn, "bot", "1", queue)
}

func TestFurbyDriver(t *testing.T) {
	var err interface{}

	d := initTestFurbyDriver(newGpioTestAdaptor("adaptor"), nil)

	gobot.Assert(t, d.Name(), "bot")
	gobot.Assert(t, d.Pin(), "1")
	gobot.Assert(t, d.Connection().Name(), "adaptor")

	testAdaptorDigitalWrite = func() (err error) {
		return errors.New("write error")
	}
	testAdaptorPwmWrite = func() (err error) {
		return errors.New("pwm error")
	}

	err = d.Command("Toggle")(nil)
	gobot.Assert(t, err.(error), errors.New("write error"))

	err = d.Command("On")(nil)
	gobot.Assert(t, err.(error), errors.New("write error"))

	err = d.Command("Off")(nil)
	gobot.Assert(t, err.(error), errors.New("write error"))

}

func TestFurbyDriverStart(t *testing.T) {
	d := initTestFurbyDriver(newGpioTestAdaptor("adaptor"), nil)
	gobot.Assert(t, len(d.Start()), 0)
}

func TestFurbyDriverHalt(t *testing.T) {
	d := initTestFurbyDriver(newGpioTestAdaptor("adaptor"), nil)
	gobot.Assert(t, len(d.Halt()), 0)
}

func TestFurbyDriverToggle(t *testing.T) {
	d := initTestFurbyDriver(newGpioTestAdaptor("adaptor"), nil)
	d.Off()
	d.Toggle()
	gobot.Assert(t, d.State(), true)
	d.Toggle()
	gobot.Assert(t, d.State(), false)
}

func TestListCommands(t *testing.T) {
	d := initTestFurbyDriver(newGpioTestAdaptor("adaptor"), nil)
	commands := d.ListCommands()
	gobot.Assert(t, len(commands), 2)
}
