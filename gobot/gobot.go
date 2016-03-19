package gobot

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
	
	"time"
)

func FurbyControl() {
	gbot := gobot.NewGobot()
	
	r := raspi.NewRaspiAdaptor("raspi")
	furby := gpio.NewLedDriver(r, "led", "7")
	
	work := func() {
		gobot.Every(1*time.Second, func() {
			furby.Toggle()
		})
	}

	robot := gobot.NewRobot("furby",
		[]gobot.Connection{r},
		[]gobot.Device{furby},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
