package Bridge

// IDevice is the interface for device operations.
type IDevice interface {
	enable()
	disable()
	isEnabled() bool
	getVolume() int64
	getChannel() int64
	setVolume(vol int64)
	setChannel(channel int64)
}

// Radio represents a radio device.
type Radio struct {
	channel int64
	volume  int64
	radioEn bool
}

// Implementing methods for the Radio.
func (r *Radio) enable() {
	r.radioEn = true
}
func (r *Radio) disable() {
	r.radioEn = false
}
func (r *Radio) isEnabled() bool {
	return r.radioEn
}
func (r *Radio) getVolume() int64 {
	return r.volume
}
func (r *Radio) getChannel() int64 {
	return r.channel
}
func (r *Radio) setVolume(vol int64) {
	r.volume = vol
}
func (r *Radio) setChannel(channel int64) {
	r.channel = channel
}

// TV represents a TV device.
type TV struct {
	channel int64
	volume  int64
	tvEn    bool
}

// Implementing methods for the TV.
func (t *TV) enable() {
	t.tvEn = true
}
func (t *TV) disable() {
	t.tvEn = false
}
func (t *TV) isEnabled() bool {
	return t.tvEn
}
func (t *TV) getVolume() int64 {
	return t.volume
}
func (t *TV) getChannel() int64 {
	return t.channel
}
func (t *TV) setVolume(vol int64) {
	t.volume = vol
}
func (t *TV) setChannel(channel int64) {
	t.channel = channel
}

// Remote represents a generic remote control.
type Remote struct {
	device IDevice
}

// TogglePower toggles the power state of the device.
func (r *Remote) togglePower() {
	if r.device.isEnabled() {
		r.device.disable()
		return
	}
	r.device.enable()
}

// VolumeDown decreases the volume of the device.
func (r *Remote) volumeDown() {
	currentVolume := r.device.getVolume()
	r.device.setVolume(currentVolume - 1)
}

// VolumeUp increases the volume of the device.
func (r *Remote) volumeUp() {
	currentVolume := r.device.getVolume()
	r.device.setVolume(currentVolume + 1)
}

// ChannelDown decreases the channel of the device.
func (r *Remote) channelDown() {
	currentChannel := r.device.getChannel()
	r.device.setChannel(currentChannel - 1)
}

// ChannelUp increases the channel of the device.
func (r *Remote) channelUp() {
	currentChannel := r.device.getChannel()
	r.device.setChannel(currentChannel + 1)
}

// AdvancedRemote represents an advanced remote control with mute functionality.
type AdvancedRemote struct {
	Remote
	isMute  bool
	prevVol int64
}

// Mute toggles the mute state of the device.
func (ar *AdvancedRemote) mute() {
	if ar.isMute {
		ar.device.setVolume(ar.prevVol)
		ar.isMute = false
		return
	}
	ar.prevVol = ar.device.getVolume()
	ar.device.setVolume(0)
	ar.isMute = true
}

// SimpleRemote represents a simple remote control.
type SimpleRemote struct {
	Remote
}
