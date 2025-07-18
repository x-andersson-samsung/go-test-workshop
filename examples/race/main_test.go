package race

import (
	"testing"
)

// Try running following commands:
// go test -count=1 .        // Should pass as it does not detect a data race issue
// go test -count=1 -race .  // Should fail
func TestRace(t *testing.T) {
	Race()
}
