package Factory_Method

import (
	"testing"
)

func TestVehicleFactoryMethod(t *testing.T) {
	carFactorCreator := &CarFactory{}
	truckFactorCreator := &TruckFactory{}
	motorcycleFactorCreator := &MotorcycleFactory{}
	// Test creating a car
	car := carFactorCreator.CreateVehicle("Toyota", "Camry", 120)
	if _, ok := car.(*Car); !ok {
		t.Errorf("Expected a Car, but got %T", car)
	}
	// Test creating a truck
	truck := truckFactorCreator.CreateVehicle("Ford", "F-150", 80.0)
	if _, ok := truck.(*Truck); !ok {
		t.Errorf("Expected a Truck, but got %T", truck)
	}
	// Test creating a motorcycle
	motorcycle := motorcycleFactorCreator.CreateVehicle("Harley-Davidson", "Sportster", 150.0)
	if _, ok := motorcycle.(*Motorcycle); !ok {
		t.Errorf("Expected a Motorcycle, but got %T", motorcycle)
	}

	// Test StartEngine and StopEngine methods
	car.StartEngine()
	if !car.(*Car).EngineRunning {
		t.Errorf("Expected car's engine to be running after StartEngine")
	}
	car.StopEngine()
	if car.(*Car).EngineRunning {
		t.Errorf("Expected car's engine to be stopped after StopEngine")
	}

	// Test GetBrand, GetModel, and GetMaxSpeed methods
	brand := car.GetBrand()
	if brand != "Toyota" {
		t.Errorf("Expected brand to be 'Toyota', but got '%s'", brand)
	}

	model := car.GetModel()
	if model != "Camry" {
		t.Errorf("Expected model to be 'Camry', but got '%s'", model)
	}

	maxSpeed := car.GetMaxSpeed()
	if maxSpeed != 120.0 {
		t.Errorf("Expected max speed to be 120.0, but got %f", maxSpeed)
	}
}
