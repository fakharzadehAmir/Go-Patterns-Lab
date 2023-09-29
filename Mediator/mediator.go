package Mediator

import "fmt"

// ITrafficControlMediator is the Mediator interface defining methods for traffic control.
type ITrafficControlMediator interface {
	canArrive(aircraft IAircraft) bool
	notifyAircraft()
}

// ATC is a concrete Mediator that manages aircraft arrivals and departures.
type ATC struct {
	allAircraft     []IAircraft
	readyForLanding bool
}

// canArrive checks if an aircraft can arrive and adds it to the queue if not ready.
func (atc *ATC) canArrive(aircraft IAircraft) bool {
	if atc.readyForLanding {
		atc.readyForLanding = false
		return true
	}
	atc.allAircraft = append(atc.allAircraft, aircraft)
	return false
}

// notifyAircraft signals the next aircraft to land and makes the runway ready.
func (atc *ATC) notifyAircraft() {
	if !atc.readyForLanding {
		atc.readyForLanding = true
	}
	if len(atc.allAircraft) > 0 {
		nextAircraft := atc.allAircraft[0]
		atc.allAircraft = atc.allAircraft[1:]
		nextAircraft.permission()
	}
}

// IAircraft is the Colleague interface representing aircraft.
type IAircraft interface {
	permission()
	departure()
	arrival()
}

// PrivateAircraft is a concrete Colleague representing private aircraft.
type PrivateAircraft struct {
	mediator ITrafficControlMediator
}

// arrival handles arrival for private aircraft.
func (pa *PrivateAircraft) arrival() {
	if !pa.mediator.canArrive(pa) {
		fmt.Println("Private Aircraft wait")
		return
	}
	fmt.Println("Private Aircraft arrived")
}

// departure handles departure for private aircraft.
func (pa *PrivateAircraft) departure() {
	fmt.Println("Private Aircraft can leave")
	pa.mediator.notifyAircraft()
}

// permission handles permission for private aircraft.
func (pa *PrivateAircraft) permission() {
	fmt.Println("Private Aircraft can arrive")
	pa.arrival()
}

// CommercialAircraft is a concrete Colleague representing commercial aircraft.
type CommercialAircraft struct {
	mediator ITrafficControlMediator
}

// arrival handles arrival for commercial aircraft.
func (ca *CommercialAircraft) arrival() {
	if !ca.mediator.canArrive(ca) {
		fmt.Println("Commercial Aircraft wait")
		return
	}
	fmt.Println("Commercial Aircraft arrived")
}

// departure handles departure for commercial aircraft.
func (ca *CommercialAircraft) departure() {
	fmt.Println("Commercial Aircraft can leave")
	ca.mediator.notifyAircraft()
}

// permission handles permission for commercial aircraft.
func (ca *CommercialAircraft) permission() {
	fmt.Println("Commercial Aircraft can arrive")
	ca.arrival()
}
