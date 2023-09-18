package Composite

import (
	"testing"
)

func TestEmployeeSalary(t *testing.T) {
	// Create a Worker
	worker := NewWorker(50000.0)

	// Test the salary of the Worker
	expectedSalary := 50000.0
	if salary := worker.GetSalary(); salary != expectedSalary {
		t.Errorf("Expected salary: %.2f, got %.2f", expectedSalary, salary)
	}
}

func TestManagerSalary(t *testing.T) {
	// Create Workers
	worker1 := NewWorker(50000.0)
	worker2 := NewWorker(60000.0)

	// Create a Manager with Workers as subordinates
	manager := NewManager(70000.0)
	manager.AddEmployee(worker1)
	manager.AddEmployee(worker2)

	// Test the salary of the Manager (including subordinates)
	expectedSalary := 180000.0 // 70000 (Manager) + 50000 (Worker1) + 60000 (Worker2)
	if salary := manager.GetSalary(); salary != expectedSalary {
		t.Errorf("Expected salary: %.2f, got %.2f", expectedSalary, salary)
	}
}

func TestManagerWithoutSubordinates(t *testing.T) {
	// Create an empty Manager
	manager := NewManager(70000.0)

	// Test the salary of the Manager (no subordinates)
	expectedSalary := 70000.0
	if salary := manager.GetSalary(); salary != expectedSalary {
		t.Errorf("Expected salary: %.2f, got %.2f", expectedSalary, salary)
	}
}
