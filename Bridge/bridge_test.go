package Bridge

import (
	"testing"
)

func TestAdvancedRemoteControl(t *testing.T) {
	// Create a radio and an advanced remote control for the radio
	radio := &Radio{}
	advancedRemote := &AdvancedRemote{Remote: Remote{device: radio}}

	// Test turning on and off the radio
	advancedRemote.togglePower()
	if !radio.isEnabled() {
		t.Error("Expected radio to be on.")
	}

	// Test muting and unmuting the radio
	radio.setVolume(10)
	advancedRemote.mute()
	if radio.getVolume() != 0 {
		t.Errorf("Expected volume to be 0, got %d", radio.getVolume())
	}
	advancedRemote.mute()
	if radio.getVolume() != 10 {
		t.Errorf("Expected volume to be 10, got %d", radio.getVolume())
	}

	// Test volume control
	radio.setVolume(5)
	advancedRemote.volumeUp()
	if radio.getVolume() != 6 {
		t.Errorf("Expected volume to be 6, got %d", radio.getVolume())
	}
	advancedRemote.volumeDown()
	if radio.getVolume() != 5 {
		t.Errorf("Expected volume to be 5, got %d", radio.getVolume())
	}

	// Test channel control (not applicable to radio, should have no effect)
	radio.setChannel(3)
	advancedRemote.channelUp()
	if radio.getChannel() != 4 {
		t.Errorf("Expected channel to be 4, got %d", radio.getChannel())
	}
	advancedRemote.channelDown()
	if radio.getChannel() != 3 {
		t.Errorf("Expected channel to be 3, got %d", radio.getChannel())
	}

	// Test turning off the radio
	advancedRemote.togglePower()
	if radio.isEnabled() {
		t.Error("Expected radio to be off.")
	}
}
func TestSimpleRemoteControl(t *testing.T) {
	// Create a TV and a simple remote control for the TV
	tv := &TV{}
	simpleRemote := &SimpleRemote{Remote: Remote{device: tv}}

	// Test turning on and off the TV
	simpleRemote.togglePower()
	if !tv.isEnabled() {
		t.Error("Expected TV to be on.")
	}
	simpleRemote.togglePower()
	if tv.isEnabled() {
		t.Error("Expected TV to be off.")
	}

	// Test volume control
	tv.setVolume(10)
	simpleRemote.volumeUp()
	if tv.getVolume() != 11 {
		t.Errorf("Expected volume to be 11, got %d", tv.getVolume())
	}
	simpleRemote.volumeDown()
	if tv.getVolume() != 10 {
		t.Errorf("Expected volume to be 10, got %d", tv.getVolume())
	}

	// Test channel control
	tv.setChannel(5)
	simpleRemote.channelUp()
	if tv.getChannel() != 6 {
		t.Errorf("Expected channel to be 6, got %d", tv.getChannel())
	}
	simpleRemote.channelDown()
	if tv.getChannel() != 5 {
		t.Errorf("Expected channel to be 5, got %d", tv.getChannel())
	}
}
