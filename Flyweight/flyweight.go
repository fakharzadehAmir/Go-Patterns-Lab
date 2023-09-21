package Flyweight

const (
	awayTeam = "away"
	homeTeam = "home"
)

// Player represents a football player.
type Player struct {
	Firstname string
	Lastname  string
	Overall   int
	LatLong   struct {
		lat  int
		long int
	}
	Jersey IJersey // The Jersey field is part of the Flyweight pattern.
}

// JerseyFactory is responsible for creating and managing jersey instances.
type JerseyFactory struct {
	jerseyMap map[string]IJersey // A map to store jersey instances.
}

// getJerseyBySide returns a shared jersey instance based on the side (home or away) and color.
// It follows the Flyweight pattern by reusing existing jersey instances when possible.
func (jf *JerseyFactory) getJerseyBySide(side, color string) IJersey {
	if jf.jerseyMap[side] != nil {
		return jf.jerseyMap[side] // Return the existing jersey instance.
	} else if side == awayTeam {
		jf.jerseyMap[side] = NewAwayJersey(color) // Create a new jersey instance if none exists.
		return jf.jerseyMap[side]
	} else if side == homeTeam {
		jf.jerseyMap[side] = NewHomeJersey(color) // Create a new jersey instance if none exists.
		return jf.jerseyMap[side]
	}
	return nil // Return nil if an invalid side is provided.
}

// NewJerseyFactory creates a new JerseyFactory instance.
func NewJerseyFactory() *JerseyFactory {
	return &JerseyFactory{jerseyMap: make(map[string]IJersey)}
}

// IJersey is the interface that represents a football jersey.
type IJersey interface {
	getColor() string
}

// HomeJersey represents a home team jersey.
type HomeJersey struct {
	color string
}

// getColor returns the color of the home jersey.
func (hj *HomeJersey) getColor() string {
	return hj.color
}

// NewHomeJersey creates a new instance of HomeJersey with the given color.
func NewHomeJersey(color string) *HomeJersey {
	return &HomeJersey{color: color}
}

// AwayJersey represents an away team jersey.
type AwayJersey struct {
	color string
}

// getColor returns the color of the away jersey.
func (aj *AwayJersey) getColor() string {
	return aj.color
}

// NewAwayJersey creates a new instance of AwayJersey with the given color.
func NewAwayJersey(color string) *AwayJersey {
	return &AwayJersey{color: color}
}
