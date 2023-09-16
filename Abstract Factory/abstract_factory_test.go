package Abstract_Factory

import "testing"

func TestManFactory(t *testing.T) {
	manFactory := &ManFactory{}

	// Create men's soccer team
	manSoccerTeam := manFactory.CreateSoccerTeam(5, 3)
	if manSoccerTeam.GetNumberOfWin() != 5 || manSoccerTeam.GetNumberOfLose() != 3 {
		t.Errorf("Expected 5 wins and 3 losses for men's soccer team, but got %d wins and %d losses.", manSoccerTeam.GetNumberOfWin(), manSoccerTeam.GetNumberOfLose())
	}

	// Create men's basketball team
	manBasketballTeam := manFactory.CreateBasketballTeam(100)
	if manBasketballTeam.GetNumberOfShots() != 100 {
		t.Errorf("Expected 100 shots for men's basketball team, but got %d shots.", manBasketballTeam.GetNumberOfShots())
	}

	// Create men's volleyball team
	manVolleyballTeam := manFactory.CreateVolleyballTeam(50)
	if manVolleyballTeam.GetNumberOfService() != 50 {
		t.Errorf("Expected 50 services for men's volleyball team, but got %d services.", manVolleyballTeam.GetNumberOfService())
	}
}

func TestWomanFactory(t *testing.T) {
	womanFactory := &WomanFactory{}

	// Create women's soccer team
	womanSoccerTeam := womanFactory.CreateSoccerTeam(4, 2)
	if womanSoccerTeam.GetNumberOfWin() != 4 || womanSoccerTeam.GetNumberOfLose() != 2 {
		t.Errorf("Expected 4 wins and 2 losses for women's soccer team, but got %d wins and %d losses.", womanSoccerTeam.GetNumberOfWin(), womanSoccerTeam.GetNumberOfLose())
	}

	// Create women's basketball team
	womanBasketballTeam := womanFactory.CreateBasketballTeam(80)
	if womanBasketballTeam.GetNumberOfShots() != 80 {
		t.Errorf("Expected 80 shots for women's basketball team, but got %d shots.", womanBasketballTeam.GetNumberOfShots())
	}

	// Create women's volleyball team
	womanVolleyballTeam := womanFactory.CreateVolleyballTeam(40)
	if womanVolleyballTeam.GetNumberOfService() != 40 {
		t.Errorf("Expected 40 services for women's volleyball team, but got %d services.", womanVolleyballTeam.GetNumberOfService())
	}
}
