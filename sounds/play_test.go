package sounds

import (
	"testing"
)

func TestNoFile(t *testing.T) {

	err := PlayFile("foo.wav")
	if err == nil {
		t.Fatal("Should have failed for a non-existent file.")
	}

}

func TestNonWavFile(t *testing.T) {
	err := PlayFile("badFile.txt")
	if err == nil {
		t.Fatal("Gave a non wav file..  should have failed.")
	}
}

func TestRealFile(t *testing.T) {
	err := PlayFile("boing.wav")
	if err != nil {
		t.Fatal("This test file from golang audio should have played.")
	}
}
