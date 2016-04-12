package jenkins

import (
	"testing"
)

func TestEmptyTest(t *testing.T) {

	t.Error("Should not be able to pass with empty test.")
}
