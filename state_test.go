// Package for testing the represention of the state of the riverworld

package state

import "testing"

// Test function implemented in line with the Golang's testing tool

func TestViewState(t *testing.T) {
	wanted := "[  ---\\    _________________\\  / / human sheep wolf corn]"
//Parameter 1 bør tar int 0-3
//Parameter 1 == 0 Man tar wolf on the boat
//Parameter 2 == 1 Man tar corn on the boat
//Parameter 1 == 2 Man krosser alene den elven
//Parameter 1 == 3 man tar sheep on the boat
// Hvis du finner right condition ,skal man put of the item on the left
//and carry to right.state.state
// For 1st item enter 0 for parameter 2
// For 2nd item enter 1 for parameter 2
// Parameter 3 is third choice, if you choose right human at logic and
// finalize the process. You should choose 0 or 1.state
// Parameter 3 = 1 means human turn right side alone.state

	state := ViewState(3, 0, 0)
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
	}
