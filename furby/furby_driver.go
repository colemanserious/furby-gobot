package furby

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"strings"
)

var _ gobot.Driver = (*FurbyDriver)(nil)

// Prepend for all files to get full location
//  Should be from execution path of binary (base)
var resourceRoot string = "resources/"

// Map command name to file name for command
var commands = map[string]string{
	"whee": "whee.wav",
	"burp": "burp.wav",
}

var commandList = []string{}

func init() {
	for k, v := range commands {
		commandList = append(commandList, k)
		commands[k] = strings.Join([]string{resourceRoot, v}, "")
	}
}

type FurbyDriver struct {
	pin        string
	name       string
	connection gpio.DigitalWriter
	high       bool
	gobot.Commander
	soundQueue chan string
}

func NewFurbyDriver(a gpio.DigitalWriter, name string, pin string, soundQueue chan string) *FurbyDriver {
	f := &FurbyDriver{
		name:       name,
		pin:        pin,
		connection: a,
		high:       false,
		soundQueue: soundQueue,
		Commander:  gobot.NewCommander(),
	}
	f.AddCommand("Toggle", func(params map[string]interface{}) interface{} {
		return f.Toggle()
	})

	f.AddCommand("On", func(params map[string]interface{}) interface{} {
		return f.On()
	})

	f.AddCommand("Off", func(params map[string]interface{}) interface{} {
		return f.Off()
	})

	return f
}

func (f *FurbyDriver) Name() string { return f.name }

func (f *FurbyDriver) Connection() gobot.Connection {
	return f.connection.(gobot.Connection)
}

func (f *FurbyDriver) Start() (errs []error) {
	return
}

func (f *FurbyDriver) Halt() (errs []error) {
	return
}

// State return true if the led is On and false if the led is Off
func (f *FurbyDriver) State() bool {
	return f.high
}

// Pin returns the GPIO pin in use
func (f *FurbyDriver) Pin() string { return f.pin }

// On sets the led to a high state.
func (f *FurbyDriver) On() (err error) {
	if err = f.connection.DigitalWrite(f.Pin(), 1); err != nil {
		return
	}
	f.high = true
	return
}

// Off sets the led to a low state.
func (f *FurbyDriver) Off() (err error) {
	if err = f.connection.DigitalWrite(f.Pin(), 0); err != nil {
		return
	}
	f.high = false
	return
}

// Toggle sets the led to the opposite of it's current state
func (f *FurbyDriver) Toggle() (err error) {
	if f.State() {
		err = f.Off()
	} else {
		err = f.On()
	}
	return
}

func (f *FurbyDriver) ExecuteCommand(command string) (err error) {

	if file, ok := commands[command]; ok {
		f.soundQueue <- file
		return
	} else {
		return fmt.Errorf("Command %v not available", command)
	}
	return

}

func (f *FurbyDriver) ListCommands() []string {
	return commandList
}
