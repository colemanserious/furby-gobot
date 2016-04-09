package sounds

import (
	"testing"
)

func TestEmptyFileName(t *testing.T) {

	err := PlayWav("")
	if err == nil {
		t.Error("Should not be able to play an empty string file.")
	}

}

func TestRealFileName(t *testing.T) {

	err := PlayWav("boing.wav")
	if err != nil {
		t.Error("Should have been able to play file.")
	}
}
