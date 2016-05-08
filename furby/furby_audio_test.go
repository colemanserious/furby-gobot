package furby

import (
	"github.com/colemanserious/furby-gobot/audio"
	"github.com/hybridgroup/gobot"
	//	"github.com/hybridgroup/gobot/platforms/gpio"
	"testing"
)

var d *FurbyDriver

func initTestInfrastructure() {
	queue := make(chan string, 1)
	d = initTestFurbyDriver(newGpioTestAdaptor("adaptor"), queue)
	a := audio.NewAudioDriver(audio.NewAudioAdaptor("conn"), "dev", queue)
	a.Start()
}

func TestFurbyKnownCommands(t *testing.T) {

	// test output should have messages indicating that each of these has a corollary sound played
	initTestInfrastructure()
	err := d.ExecuteCommand("fart")
	gobot.Assert(t, err, nil)
	err = d.ExecuteCommand("laugh")
	gobot.Assert(t, err, nil)
	err = d.ExecuteCommand("burp")
	gobot.Assert(t, err, nil)
}

func TestFurbyUnknownCommand(t *testing.T) {
	// test output should have no message indicating that a sound was played
	initTestInfrastructure()
	err := d.ExecuteCommand("unknown")
	// should receive an error for an unknown command
	gobot.Assert(t, err != nil, true)
}
