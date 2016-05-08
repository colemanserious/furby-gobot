package audio

import (
	"os"
	"errors"
	"github.com/hybridgroup/gobot"
	"log"
	"os/exec"
)

var _ gobot.Adaptor = (*AudioAdaptor)(nil)

type AudioAdaptor struct {
	name string
}

func NewAudioAdaptor(name string) *AudioAdaptor {
	return &AudioAdaptor{
		name: name,
	}
}

func (a *AudioAdaptor) Name() string { return a.name }

func (a *AudioAdaptor) Connect() []error { return nil }

func (a *AudioAdaptor) Finalize() []error { return nil }

func (a *AudioAdaptor) Sound(fileName string) []error {

	var errorsList []error
	var err error

	if fileName == "" {
		log.Println("Require filename for WAV file.")
		errorsList = append(errorsList, errors.New("Requires filename for WAV file."))
		return errorsList
	}

	_, err = os.Open(fileName)
	if err != nil {
		log.Println(err)
		errorsList = append(errorsList, err)
		return errorsList
	}

	// command to play a WAV file on Raspberry Pi
	cmd := exec.Command("aplay", fileName)
	err = cmd.Run()

	if err != nil {
		log.Println(err)
		errorsList = append(errorsList, err)
		return errorsList
	}

	// Need to return to fulfill function sig, even though returning an empty
	return nil
}
