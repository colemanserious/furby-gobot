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
	"github.com/colemanserious/furby-gobot/jenkinsconnect"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"github.com/spf13/cobra"
)

// listenJenkinsCmd represents the listenJenkins command
var listenJenkinsCmd = &cobra.Command{
	Use:   "listenJenkins",
	Short: "Sets up endpoint Jenkins can use to notify job results",
	Long: `Sets up an http endpoint on localhost:3000.  That endpoint includes a command 'parseResult', which will receive a data payload, determine the outcome of the Jenkins job, and fire off handlers for appropriate robot notifications.
	
	For this incarnation, we'll be driving the Furby, as well as updating an LCD, all by invoking Go interfaces provided by other packages, brought together via Gobot

`,
	Run: func(cmd *cobra.Command, args []string) {
		gbot := gobot.NewGobot()
		api.NewAPI(gbot).Start()

		r := raspi.NewRaspiAdaptor("raspi")
		jenkinsConnect := jenkinsconnect.NewJenkinsconnectAdaptor("jenkins")

		jenkinsDriver := jenkinsconnect.NewJenkinsconnectDriver(jenkinsConnect, "jenkins-command")

		work := func() {
			gobot.On(jenkinsDriver.Event("jobResult"), func(data interface{}) {
				fmt.Println("Received jobResult event")
			})

		}

		robot := gobot.NewRobot("furbyBot",
			[]gobot.Connection{r},
			[]gobot.Device{jenkinsDriver},
			work,
		)

		gbot.AddRobot(robot)
		gbot.Start()

	},
}

func init() {
	RootCmd.AddCommand(listenJenkinsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenJenkinsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenJenkinsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
