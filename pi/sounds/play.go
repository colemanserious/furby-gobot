package sounds

import (
	"errors"
	"log"
	"os/exec"
	// if don't comment out below, go test will complain, and the software won't build
	//"os"
)

// PlayWav plays the filename provided.
//  Potential errors include:
//   - file not found
//   - unable to play
//
// Note: Did see audio package in GoLang, however its part of gomobile
//  Other alternatives exist - OpenAL, portaudio, ...  aplay seemed simplest
func PlayWav(fileName string) (err error) {

	if fileName == "" {
		log.Println("Require filename for WAV file.")
		return errors.New("Requires filename for WAV file.")
	}

	// command to play a WAV file on Raspberry Pi
	cmd := exec.Command("aplay", fileName)
	err = cmd.Run()

	if err != nil {
		log.Println(err)
		return err
	}

	// Need to return to fulfill function sig, even though returning an empty
	return
}
