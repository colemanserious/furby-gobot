// Copyright © 2016 Tina Coleman <colemanserious@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/colemanserious/furby-gobot/audio"
	"github.com/colemanserious/furby-gobot/jenkinsconnect"
	"github.com/hybridgroup/gobot/api"
	"github.com/spf13/cobra"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"github.com/hybridgroup/gobot/platforms/raspi"

	"log"
	"time"
)

// ledOnCmd represents the ledOn command
var ledOnCmd = &cobra.Command{
	Use:   "ledOn",
	Short: "Toggle led on pin 13",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ledOn called")

		gbot := gobot.NewGobot()
		api.NewAPI(gbot).Start()

		r := raspi.NewRaspiAdaptor("raspi")
		audioAdaptor := audio.NewAudioAdaptor("sound")
		jenkinsConnect := jenkinsconnect.NewJenkinsconnectAdaptor("jenkins")

		led := gpio.NewLedDriver(r, "led", "13")
		audioDriver := audio.NewAudioDriver(audioAdaptor, "sounds", nil)
		jenkinsDriver := jenkinsconnect.NewJenkinsconnectDriver(jenkinsConnect, "jenkins-command")

		screen := i2c.NewGroveLcdDriver(r, "screen")
		work := func() {

			screen.Clear()
			screen.Home()

			//screen.SetRGB(255, 0, 0)

			gobot.Every(5*time.Second, func() {
				led.Toggle()
				audioDriver.Sound("resources/foo.wav")
				if err := screen.Write("Writing, writing..."); err != nil {
					log.Fatal(err)
				}

				screen.SetRGB(0, 255, 0)
				// set a custom character in the first position
				screen.SetCustomChar(0, i2c.CustomLCDChars["smiley"])
				// add the custom character at the end of the string
				screen.Write("goodbye\nhave a nice day " + string(byte(0)))
				gobot.Every(500*time.Millisecond, func() {
					screen.Scroll(false)
				})
			})

			<-time.After(1 * time.Second)
			//screen.SetRGB(0, 0, 255)
		}

		robot := gobot.NewRobot("furbyBot",
			[]gobot.Connection{r, audioAdaptor},
			[]gobot.Device{led, audioDriver, screen, jenkinsDriver},
			work,
		)

		gbot.AddRobot(robot)

		gbot.Start()
	},
}

func init() {
	RootCmd.AddCommand(ledOnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ledOnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ledOnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
