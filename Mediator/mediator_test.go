package Mediator

import "testing"

func TestTrafficControl(t *testing.T) {
	// Create an ATC mediator
	atc := &ATC{}

	// Create private and commercial aircraft
	privateAircraft := &PrivateAircraft{mediator: atc}
	commercialAircraft := &CommercialAircraft{mediator: atc}

	// Test private aircraft arrival and departure
	privateAircraft.arrival()
	if len(atc.allAircraft) == 0 {
		t.Errorf("Expected 0 aircraft in queue, got %d", len(atc.allAircraft))
	}

	privateAircraft.departure()
	if atc.readyForLanding {
		t.Error("Expected runway to be ready after private aircraft departure")
	}

	// Test commercial aircraft arrival
	commercialAircraft.arrival()
	if atc.readyForLanding {
		t.Error("Expected runway not to be ready for commercial aircraft arrival")
	}

	// Test commercial aircraft departure
	commercialAircraft.departure()
	if atc.readyForLanding {
		t.Error("Expected runway to be ready after commercial aircraft departure")
	}
	if len(atc.allAircraft) != 0 {
		t.Errorf("Expected 0 aircraft in queue, got %d", len(atc.allAircraft))
	}
}
