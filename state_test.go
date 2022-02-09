// Package for testing the represention of the state of the riverworld

package state

import "testing"

// Test function implemented in line with the Golang's testing tool

func TestViewState(t *testing.T) {
	wanted := "[  ---\\    _________________\\  / / human sheep wolf corn]"

	state := ViewState(3, 1, 1)
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
	}
