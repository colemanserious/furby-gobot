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
	"github.com/colemanserious/furby-gobot/furby"
	"github.com/colemanserious/furby-gobot/jenkinsconnect"
	"github.com/hybridgroup/gobot/api"
	"github.com/spf13/cobra"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"github.com/hybridgroup/gobot/platforms/raspi"

	//"time"
	"log"
)

// furbyBotCmd represents the furbyBot command
var furbyBotCmd = &cobra.Command{
	Use:   "furbyBot",
	Short: "Primary command for project - set up gobot Furby interaction",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("furbyBot called")

		gbot := gobot.NewGobot()
		api.NewAPI(gbot).Start()

		r := raspi.NewRaspiAdaptor("raspi")
		audioAdaptor := audio.NewAudioAdaptor("sound")
		jenkinsAdaptor := jenkinsconnect.NewJenkinsconnectAdaptor("jenkins")

		// Set up asynchronous channel - if we get more than 3 sounds being played in a row, something's up
		csoundFiles := make(chan string, 3)
		furby := furby.NewFurbyDriver(r, "furby", "16", csoundFiles)
		audioDriver := audio.NewAudioDriver(audioAdaptor, "sounds", csoundFiles)
		jenkinsDriver := jenkinsconnect.NewJenkinsconnectDriver(jenkinsAdaptor, "jenkins-command")

		screen := i2c.NewGroveLcdDriver(r, "screen")
		work := func() {

			screen.Clear()
			screen.Home()
			furby.On()

			screen.SetRGB(255, 255, 255)

			if err := screen.Write("Furby say hi!!"); err != nil {
				log.Fatal(err)
			}

			gobot.On(jenkinsDriver.Event("jobResult"), func(data interface{}) {
				jobResult := data.(jenkinsconnect.JobOutcome)
				switch jobResult.State {
				case jenkinsconnect.SUCCESS:
					screen.Home()
					screen.Clear()
					screen.SetRGB(0, 255, 0)
					screen.Write("Success:\n" + jobResult.Name)
					furby.ExecuteCommand("burp")
				case jenkinsconnect.FAILED:
					screen.Home()
					screen.Clear()
					screen.SetRGB(255, 0, 0)
					screen.Write(" FAIL:\n" + jobResult.Name)
					furby.ExecuteCommand("fart")
				default:
				}
			})

		}

		robot := gobot.NewRobot("furbyBot",
			[]gobot.Connection{r, audioAdaptor},
			[]gobot.Device{furby, audioDriver, screen, jenkinsDriver},
			//[]gobot.Device{furby, audioDriver, jenkinsDriver},
			work,
		)

		gbot.AddRobot(robot)

		gbot.Start()
	},
}

func init() {
	RootCmd.AddCommand(furbyBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ledOnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ledOnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
