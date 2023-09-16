package Abstract_Factory

// ISoccerTeam is an interface for soccer teams.
type ISoccerTeam interface {
	GetNumberOfWin() int64
	GetNumberOfLose() int64
}

// IBasketballTeam is an interface for basketball teams.
type IBasketballTeam interface {
	GetNumberOfShots() int64
}

// IVolleyballTeam is an interface for volleyball teams.
type IVolleyballTeam interface {
	GetNumberOfService() int64
}

// ManSoccerTeam represents a men's soccer team.
type ManSoccerTeam struct {
	Win  int64
	Lose int64
}

// NewManTeamSoccer creates a new instance of ManSoccerTeam.
func NewManTeamSoccer(win, lose int64) ISoccerTeam {
	return &ManSoccerTeam{
		Win:  win,
		Lose: lose,
	}
}

// GetNumberOfWin returns the number of wins for the men's soccer team.
func (ms *ManSoccerTeam) GetNumberOfWin() int64 {
	return ms.Win
}

// GetNumberOfLose returns the number of losses for the men's soccer team.
func (ms *ManSoccerTeam) GetNumberOfLose() int64 {
	return ms.Lose
}

// WomanSoccerTeam represents a women's soccer team.
type WomanSoccerTeam struct {
	Win  int64
	Lose int64
}

// NewWomanTeamSoccer creates a new instance of WomanSoccerTeam.
func NewWomanTeamSoccer(win, lose int64) ISoccerTeam {
	return &WomanSoccerTeam{
		Win:  win,
		Lose: lose,
	}
}

// GetNumberOfWin returns the number of wins for the women's soccer team.
func (ws *WomanSoccerTeam) GetNumberOfWin() int64 {
	return ws.Win
}

// GetNumberOfLose returns the number of losses for the women's soccer team.
func (ws *WomanSoccerTeam) GetNumberOfLose() int64 {
	return ws.Lose
}

// ManBasketballTeam represents a men's basketball team.
type ManBasketballTeam struct {
	Shots int64
}

// NewManTeamBasketball creates a new instance of ManBasketballTeam.
func NewManTeamBasketball(shots int64) IBasketballTeam {
	return &ManBasketballTeam{
		Shots: shots,
	}
}

// GetNumberOfShots returns the number of shots for the men's basketball team.
func (mbt *ManBasketballTeam) GetNumberOfShots() int64 {
	return mbt.Shots
}

// WomanBasketballTeam represents a women's basketball team.
type WomanBasketballTeam struct {
	Shots int64
}

// NewWomanTeamBasketball creates a new instance of WomanBasketballTeam.
func NewWomanTeamBasketball(shots int64) IBasketballTeam {
	return &WomanBasketballTeam{
		Shots: shots,
	}
}

// GetNumberOfShots returns the number of shots for the women's basketball team.
func (wbt *WomanBasketballTeam) GetNumberOfShots() int64 {
	return wbt.Shots
}

// ManVolleyballTeam represents a men's volleyball team.
type ManVolleyballTeam struct {
	Services int64
}

// NewManTeamVolleyball creates a new instance of ManVolleyballTeam.
func NewManTeamVolleyball(services int64) IVolleyballTeam {
	return &ManVolleyballTeam{
		Services: services,
	}
}

// GetNumberOfService returns the number of services for the men's volleyball team.
func (mvt *ManVolleyballTeam) GetNumberOfService() int64 {
	return mvt.Services
}

// WomanVolleyballTeam represents a women's volleyball team.
type WomanVolleyballTeam struct {
	Services int64
}

// NewWomanTeamVolleyball creates a new instance of WomanVolleyballTeam.
func NewWomanTeamVolleyball(services int64) IVolleyballTeam {
	return &WomanVolleyballTeam{
		Services: services,
	}
}

// GetNumberOfService returns the number of services for the women's volleyball team.
func (wvt *WomanVolleyballTeam) GetNumberOfService() int64 {
	return wvt.Services
}

// ITeamFactory is the interface for the abstract factory.
type ITeamFactory interface {
	CreateSoccerTeam(win, lose int64) ISoccerTeam
	CreateBasketballTeam(shots int64) IBasketballTeam
	CreateVolleyballTeam(service int64) IVolleyballTeam
}

// ManFactory is a concrete factory for creating men's teams.
type ManFactory struct{}

// CreateSoccerTeam creates a men's soccer team.
func (mf *ManFactory) CreateSoccerTeam(win, lose int64) ISoccerTeam {
	return NewManTeamSoccer(win, lose)
}

// CreateBasketballTeam creates a men's basketball team.
func (mf *ManFactory) CreateBasketballTeam(shots int64) IBasketballTeam {
	return NewManTeamBasketball(shots)
}

// CreateVolleyballTeam creates a men's volleyball team.
func (mf *ManFactory) CreateVolleyballTeam(service int64) IVolleyballTeam {
	return NewManTeamVolleyball(service)
}

// WomanFactory is a concrete factory for creating women's teams.
type WomanFactory struct{}

// CreateSoccerTeam creates a women's soccer team.
func (wf *WomanFactory) CreateSoccerTeam(win, lose int64) ISoccerTeam {
	return NewWomanTeamSoccer(win, lose)
}

// CreateBasketballTeam creates a women's basketball team.
func (wf *WomanFactory) CreateBasketballTeam(shots int64) IBasketballTeam {
	return NewWomanTeamBasketball(shots)
}

// CreateVolleyballTeam creates a women's volleyball team.
func (wf *WomanFactory) CreateVolleyballTeam(service int64) IVolleyballTeam {
	return NewWomanTeamVolleyball(service)
}
