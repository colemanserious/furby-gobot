package furby

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"os"
	"strings"
	"time"
)

var _ gobot.Driver = (*FurbyDriver)(nil)

var resourceRoot string

// Map command name to file name for command
//  Three digits correspond to code from Hacksby library - see Command.pm
var commands = map[string]string{
	"burp":    "864-burp.wav",
	"laugh":   "863-laugh.wav",
	"fart":    "865-fart.wav",
	"whisper": "719-whisper.wav",
	"sleep":   "718-sleep.wav",
	"listen":  "820-listen.wav",
}

var commandList = []string{}

func init() {
	// Prepend for all files to get full location
	//  Should be from execution path of binary (base)
	//  Note: dealing with '..' for case of tests...
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory: " + cwd)
	if strings.HasSuffix(cwd, "furby-gobot") {
		resourceRoot = cwd + "/resources/"
	} else {
		resourceRoot = cwd + "/../resources/"
	}

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
	go f.keepListening()

	fmt.Println("Started Furby driver")
	return
}

// keep furby "listening" and awake.  Keeps it from uttering silly things
// other than the silly ones we want it to give..
func (f *FurbyDriver) keepListening() {
	f.ExecuteCommand("listen")
	gobot.Every(40*time.Second, func() {
		f.ExecuteCommand("listen")
	})
}

func (f *FurbyDriver) Halt() (errs []error) {
	return
}

// State return true if the Furby is On and false if the Furby is Off
func (f *FurbyDriver) State() bool {
	return f.high
}

// Pin returns the GPIO pin in use
func (f *FurbyDriver) Pin() string { return f.pin }

// On sets the Furby to a high state.
func (f *FurbyDriver) On() (err error) {
	fmt.Println("Turning furby on ")
	if err = f.connection.DigitalWrite(f.Pin(), 1); err != nil {
		fmt.Println("Unable to turn furby on: %v", err)
		return err
	}
	f.high = true
	return
}

// Off sets the Furby to a low state (low = off).
func (f *FurbyDriver) Off() (err error) {
	fmt.Println("Turning furby off ")
	if err = f.connection.DigitalWrite(f.Pin(), 0); err != nil {
		fmt.Println("Unable to turn furby off %v", err)
		return
	}
	f.high = false
	return
}

// Toggle sets the Furby to the opposite of its current state
func (f *FurbyDriver) Toggle() (err error) {
	if f.State() {
		err = f.Off()
	} else {
		err = f.On()
	}
	return
}

func (f *FurbyDriver) ExecuteCommand(command string) (err error) {

	if !f.State() {
		f.On()
	}

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
