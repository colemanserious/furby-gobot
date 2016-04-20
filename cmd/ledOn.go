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
	"github.com/spf13/cobra"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"time"
)

// ledOnCmd represents the ledOn command
var ledOnCmd = &cobra.Command{
	Use:   "ledOn",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ledOn called")

		gbot := gobot.NewGobot()

		r := raspi.NewRaspiAdaptor("raspi")
		audioAdaptor := audio.NewAudioAdaptor("sound")

		led := gpio.NewLedDriver(r, "led", "13")
		audioDriver := audio.NewAudioDriver(audioAdaptor, "sounds")

		work := func() {
			gobot.Every(3*time.Second, func() {
				led.Toggle()
				audioDriver.Sound("foo.wav")
			})
		}

		robot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{r, audioAdaptor},
			[]gobot.Device{led, audioDriver},
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
