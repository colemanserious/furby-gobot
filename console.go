package main

import "github.com/colemanserious/furby-gobot/cmd"

func main() {
	if err: = cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
