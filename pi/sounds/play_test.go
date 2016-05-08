package sounds

import (
	"os/exec"
	"testing"
)

func TestEmptyFileName(t *testing.T) {

	err := PlayWav("")
	if err == nil {
		t.Error("Should not be able to play an empty string file.")
	}

}

func TestRealFileName(t *testing.T) {

	_, err := exec.LookPath("aplayer")
	if err != nil {
		t.Skip("Unable to execute command to play file.")
		return
	}
	err = PlayWav("../resources/match2.wav")
	if err != nil {
		t.Error("Should have been able to play file.", err)
	}
}
